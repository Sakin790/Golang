package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Create a mux and register routes (like main.go)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthCheck)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "OK" {
		t.Errorf("expected body OK, got %s", string(body))
	}
}


func TestGetUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUser)
	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got []User
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(got) != len(users) {
		t.Errorf("expected %d users, got %d", len(users), len(got))
	}

	for i, u := range users {
		if got[i].Name != u.Name || got[i].Email != u.Email {
			t.Errorf("expected user %+v, got %+v", u, got[i])
		}
	}
}
