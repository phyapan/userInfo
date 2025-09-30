package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

func main() {
	// Update with your MySQL credentials
	// user:password@tcp(localhost:3306)/userInfo
	dsn := "root:password@tcp(127.0.0.1:3306)/userInfo"
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

		_, err := db.Exec("INSERT INTO users (name, age, email, contact) VALUES (?, ?, ?, ?)", name, age, email, contact)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprintf(w, "Saved!")
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
