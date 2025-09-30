package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//

//go get github.com/jmoiron/sqlx
//go get github.com/lib/pq

type Product struct {
	Id          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	ImageURL    string `db:"image_url" json:"image_url"`
}

var db *sqlx.DB

func initDB() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=root dbname=testdb sslmode=disable"
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
}

func status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "API SERVER IS WORKING.......")
}

// GET all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var products []Product
	err := db.Select(&products, "SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CREATE product
func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := db.QueryRow(
		"INSERT INTO products (title, description, image_url) VALUES ($1,$2,$3) RETURNING id",
		p.Title, p.Description, p.ImageURL,
	).Scan(&p.Id)

	if err != nil {
		http.Error(w, "DB insert error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// UPDATE product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"UPDATE products SET title=$1, description=$2, image_url=$3 WHERE id=$4",
		p.Title, p.Description, p.ImageURL, id,
	)
	if err != nil {
		http.Error(w, "DB update error", http.StatusInternalServerError)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	p.Id = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// DELETE product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		http.Error(w, "DB delete error", http.StatusInternalServerError)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Product deleted")
}

func main() {
	initDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/status", status)
	mux.HandleFunc("/products", getProducts)         // GET
	mux.HandleFunc("/product/create", createProduct) // POST
	mux.HandleFunc("/product/update", updateProduct) // PUT ?id=1
	mux.HandleFunc("/product/delete", deleteProduct) // DELETE ?id=1

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
