package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Init DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate schema
	db.AutoMigrate(&User{})

	mux := http.NewServeMux()

	// Create user
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

	// Get single user by id
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Path[len("/users/"):]
		var user User

		switch r.Method {
		case http.MethodGet:
			if result := db.First(&user, id); result.Error != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodPut:
			if result := db.First(&user, id); result.Error != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			var updated User
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user.Name = updated.Name
			user.Email = updated.Email
			db.Save(&user)
			json.NewEncoder(w).Encode(user)

		case http.MethodDelete:
			if result := db.Delete(&User{}, id); result.RowsAffected == 0 {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
