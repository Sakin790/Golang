package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	db             *sql.DB
	userCache      []User
	cacheMutex     sync.RWMutex
	cacheTTL       = 10 * time.Second
	cacheTimestamp time.Time

	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "http_requests_total", Help: "Total number of HTTP requests"},
		[]string{"method", "route", "status"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "http_request_duration_seconds", Help: "Duration of HTTP requests", Buckets: prometheus.DefBuckets},
		[]string{"method", "route"},
	)
)

func main() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/k6"
	db, err = sql.Open("mysql", dsn)
	if err != nil || db.Ping() != nil {
		log.Fatal("DB Connection Failed:", err)
	}
	fmt.Println("Connected to MySQL ✅")

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Register metrics
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)

	// Routes
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", withMetrics(health, "/health"))
	http.HandleFunc("/users", withMetrics(getUsers, "/users"))
	http.HandleFunc("/user/create", withMetrics(createUser, "/user/create"))
	http.HandleFunc("/user/get/", withMetrics(getUser, "/user/get"))
	http.HandleFunc("/user/update/", withMetrics(updateUser, "/user/update"))
	http.HandleFunc("/user/delete/", withMetrics(deleteUser, "/user/delete"))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ---- Middleware for Prometheus ----
func withMetrics(handler http.HandlerFunc, route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		handler(rw, r)
		httpRequestsTotal.WithLabelValues(r.Method, route, fmt.Sprint(rw.status)).Inc()
		httpRequestDuration.WithLabelValues(r.Method, route).Observe(time.Since(start).Seconds())
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// ---- CRUD Handlers ----
func health(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }

func getUsers(w http.ResponseWriter, r *http.Request) {
	cacheMutex.RLock()
	if time.Since(cacheTimestamp) < cacheTTL && len(userCache) > 0 {
		json.NewEncoder(w).Encode(userCache)
		cacheMutex.RUnlock()
		return
	}
	cacheMutex.RUnlock()
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err == nil {
			users = append(users, u)
		}
	}
	cacheMutex.Lock()
	userCache, cacheTimestamp = users, time.Now()
	cacheMutex.Unlock()
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/get/"):]
	var u User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id=?", id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", 404)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	id, _ := result.LastInsertId()
	u.ID = int(id)
	cacheMutex.Lock()
	cacheTimestamp = time.Time{}
	cacheMutex.Unlock()
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/update/"):]
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_, err := db.Exec("UPDATE users SET name=?, email=? WHERE id=?", u.Name, u.Email, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Sscan(id, &u.ID)
	cacheMutex.Lock()
	cacheTimestamp = time.Time{}
	cacheMutex.Unlock()
	json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/delete/"):]
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	cacheMutex.Lock()
	cacheTimestamp = time.Time{}
	cacheMutex.Unlock()
	w.Write([]byte("User deleted"))
}
