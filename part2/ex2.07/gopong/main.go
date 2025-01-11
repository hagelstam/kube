package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	port := 3001

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	name := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || host == "" || name == "" {
		log.Fatal("missing required environment variables")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, name)
	fmt.Printf("connecting to: %s\n", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	db.Exec(`
		CREATE TABLE IF NOT EXISTS counter (
			count INTEGER
		)
	`)

	db.Exec(`
		INSERT INTO counter (count) 
		SELECT 0 
		WHERE NOT EXISTS (SELECT * FROM counter)
	`)

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var count uint32
		err := db.QueryRow(`
			UPDATE counter 
			SET count = count + 1 
			RETURNING count
		`).Scan(&count)
		if err != nil {
			fmt.Printf("error incrementing counter: %v\n", err)
			http.Error(w, "failed to increment counter", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Ping / Pongs: %d", count-1)))
	})

	fmt.Printf("server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
