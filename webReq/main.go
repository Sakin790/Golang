package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// apiResponse ensures every response is JSON.
type apiResponse struct {
	OK    bool        `json:"ok"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	_ = enc.Encode(v)
}

func writeError(w http.ResponseWriter, status int, err error) {
	if err == nil {
		err = errors.New(http.StatusText(status))
	}
	writeJSON(w, status, apiResponse{OK: false, Error: err.Error()})
}

// ---- In-memory model (thread-safe) ----

type Note struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type store struct {
	mu     sync.RWMutex
	nextID int64
	data   map[int64]Note
}

func newStore() *store {
	return &store{nextID: 1, data: make(map[int64]Note)}
}

func (s *store) list() []Note {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Note, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}

func (s *store) get(id int64) (Note, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[id]
	return v, ok
}

func (s *store) create(n Note) Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	n.ID = s.nextID
	s.nextID++
	s.data[n.ID] = n
	return n
}

func (s *store) update(id int64, upd Note) (Note, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	orig, ok := s.data[id]
	if !ok {
		return Note{}, false
	}
	// only Title/Done are updatable
	orig.Title = upd.Title
	orig.Done = upd.Done
	s.data[id] = orig
	return orig, true
}

func (s *store) delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[id]; !ok {
		return false
	}
	delete(s.data, id)
	return true
}

// ---- HTTP routing (pure net/http) ----

type app struct {
	store *store
	mux   http.Handler
}

func newApp() *app {
	a := &app{store: newStore()}
	// single entrypoint handler to keep 404/405 fully JSON
	a.mux = http.HandlerFunc(a.route)
	return a
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Force JSON for all replies
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Simple access log
	start := time.Now()
	a.mux.ServeHTTP(w, r)
	log.Printf("%s %s %dms", r.Method, r.URL.Path, time.Since(start).Milliseconds())
}

func (a *app) route(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/health":
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method))
			return
		}
		writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: map[string]string{"status": "ok"}})
		return

	case r.URL.Path == "/notes":
		a.handleNotesCollection(w, r)
		return

	case strings.HasPrefix(r.URL.Path, "/notes/"):
		a.handleNotesItem(w, r)
		return

	case r.URL.Path == "/debug/goroutine":
		// demo endpoint: show current goroutine ID
		writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: map[string]int{"goroutine": curGoroutineID()}})
		return
	}
	writeError(w, http.StatusNotFound, fmt.Errorf("route not found"))
}

// ---- Debug util to get current goroutine ID ----
// Go doesn’t officially expose goroutine IDs, but runtime.Stack trick works.
func curGoroutineID() int {
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	// buf looks like "goroutine 123 [running]:..."
	fields := strings.Fields(string(buf[:n]))
	if len(fields) >= 2 {
		id, _ := strconv.Atoi(fields[1])
		return id
	}
	return -1
}

func (a *app) handleNotesCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list := a.store.list()
		writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: list})
		return
	case http.MethodPost:
		var in struct {
			Title string `json:"title"`
			Done  bool   `json:"done"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			writeError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body: %w", err))
			return
		}
		in.Title = strings.TrimSpace(in.Title)
		if in.Title == "" {
			writeError(w, http.StatusBadRequest, errors.New("title is required"))
			return
		}
		created := a.store.create(Note{Title: in.Title, Done: in.Done})
		writeJSON(w, http.StatusCreated, apiResponse{OK: true, Data: created})
		return
	default:
		writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method))
	}
}

func (a *app) handleNotesItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	switch r.Method {
	case http.MethodGet:
		if note, ok := a.store.get(id); ok {
			writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: note})
			return
		}
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return

	case http.MethodPut:
		var in struct {
			Title string `json:"title"`
			Done  bool   `json:"done"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			writeError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body: %w", err))
			return
		}
		in.Title = strings.TrimSpace(in.Title)
		if in.Title == "" {
			writeError(w, http.StatusBadRequest, errors.New("title is required"))
			return
		}
		if updated, ok := a.store.update(id, Note{Title: in.Title, Done: in.Done}); ok {
			writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: updated})
			return
		}
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return

	case http.MethodDelete:
		if ok := a.store.delete(id); ok {
			writeJSON(w, http.StatusOK, apiResponse{OK: true, Data: map[string]bool{"deleted": true}})
			return
		}
		writeError(w, http.StatusNotFound, errors.New("not found"))
		return

	default:
		writeError(w, http.StatusMethodNotAllowed, fmt.Errorf("method %s not allowed", r.Method))
		return
	}
}

func main() {
	addr := ":8080"
	app := newApp()

	server := &http.Server{
		Addr:              addr,
		Handler:           app,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start server
	errCh := make(chan error, 1)
	go func() {
		log.Printf("listening on http://localhost%s", addr)
		errCh <- server.ListenAndServe()
	}()

	// OS signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("signal received: %s — shutting down...", sig)
	case err := <-errCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}

	// Give in-flight requests up to 10 seconds to finish.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
	log.Printf("server stopped")
}
