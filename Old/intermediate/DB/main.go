package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Person represents a row in the person table
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func main() {
	// Connect to MySQL
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test connection
	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Define routes
	http.HandleFunc("/persons", getPersons)          // GET all
	http.HandleFunc("/persons/add", addPerson)       // POST new
	http.HandleFunc("/persons/update", updatePerson) // PUT update
	http.HandleFunc("/persons/delete", deletePerson) // DELETE

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GET /persons
func getPersons(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT person_id, first_name FROM person")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.Name) // ✅ exactly 2 columns
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		persons = append(persons, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

// POST /persons/add?name=Sakin&age=23
func addPerson(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")
	age, _ := strconv.Atoi(ageStr)

	res, err := db.Exec("INSERT INTO person(name, age) VALUES(?, ?)", name, age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	fmt.Fprintf(w, "Inserted person with ID %d\n", id)
}

// PUT /persons/update?id=1&name=Alice&age=30
func updatePerson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")

	id, _ := strconv.Atoi(idStr)
	age, _ := strconv.Atoi(ageStr)

	_, err := db.Exec("UPDATE person SET name=?, age=? WHERE id=?", name, age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Updated person with ID %d\n", id)
}

// DELETE /persons/delete?id=1
func deletePerson(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_, err := db.Exec("DELETE FROM person WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted person with ID %d\n", id)
}
