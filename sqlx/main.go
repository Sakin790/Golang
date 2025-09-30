package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// ================== DATABASE SETUP ==================

// User model (represents table "users")
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// Global DB connection (using sqlx)
var db *sqlx.DB

// ================== HANDLERS ==================

// Create new user (INSERT)
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// Decode request body (JSON -> Go struct)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into database
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get inserted ID
	id, _ := result.LastInsertId()
	user.ID = int(id)

	// Return created user as JSON
	json.NewEncoder(w).Encode(user)
}

// Read all users (SELECT)
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	// Select all rows into users slice
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// Update user (UPDATE)
func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update by ID
	_, err := db.Exec("UPDATE users SET name=?, email=? WHERE id=?", user.Name, user.Email, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Delete user (DELETE)
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete by ID
	_, err := db.Exec("DELETE FROM users WHERE id=?", input.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

// ================== MAIN ==================
func main() {
	var err error

	// 1. Connect to SQLite database (file: "test.db")
	db, err = sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}

	// 2. Create table if not exists
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	db.MustExec(schema)

	// 3. Setup routes
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)

	fmt.Println("🚀 Server running on http://localhost:8080")
	// 4. Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
