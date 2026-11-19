package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db
}

func setupTestServer(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	// Create user + List users
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			db.Create(&user)
			json.NewEncoder(w).Encode(user)
			return
		}
		if r.Method == http.MethodGet {
			var users []User
			db.Find(&users)
			json.NewEncoder(w).Encode(users)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	return mux
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	ts := httptest.NewServer(setupTestServer(db))
	defer ts.Close()

	// Prepare JSON body
	body := []byte(`{"name":"Alice","email":"alice@example.com"}`)

	resp, err := http.Post(ts.URL+"/users", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal("Failed to decode response:", err)
	}

	if user.Name != "Alice" {
		t.Errorf("expected name Alice, got %s", user.Name)
	}
}

func TestListUsers(t *testing.T) {
	db := setupTestDB()
	db.Create(&User{Name: "Bob", Email: "bob@example.com"})
	ts := httptest.NewServer(setupTestServer(db))
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatal("Failed to decode response:", err)
	}

	if len(users) == 0 {
		t.Error("expected at least one user, got 0")
	}
}
