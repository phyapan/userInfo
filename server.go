package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

func main() {
	// Get MySQL credentials from environment variables
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DB")
	if user == "" { user = "root" }
	if pass == "" { pass = "password" }
	if host == "" { host = "127.0.0.1:3306" }
	if dbname == "" { dbname = "userInfo" }
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		age INT,
		email VARCHAR(100),
		contact VARCHAR(100)
	)`)
	if err != nil {
		log.Fatal(err)
	}

		http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			name := r.FormValue("name")
			age := r.FormValue("age")
			email := r.FormValue("email")
			contact := r.FormValue("contact")

			// Basic input validation
			if name == "" || age == "" || email == "" || contact == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "All fields are required.")
				return
			}
			if len(name) > 100 || len(email) > 100 || len(contact) > 100 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Input too long.")
				return
			}
			// Simple email format check
			if !isValidEmail(email) {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid email format.")
				return
			}

			_, err := db.Exec("INSERT INTO users (name, age, email, contact) VALUES (?, ?, ?, ?)", name, age, email, contact)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			fmt.Fprintf(w, "Saved!")
		})

		// Simple email validation function
	}

	func isValidEmail(email string) bool {
		if len(email) < 3 || len(email) > 254 {
			return false
		}
		for i := 0; i < len(email); i++ {
			if email[i] == '@' {
				return true
			}
		}
		return false
	}

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
