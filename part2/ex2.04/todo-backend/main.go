package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var todos = []Todo{}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/todos", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"todos": todos})
		}

		if r.Method == http.MethodPost {
			var newTodo Todo
			if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
				http.Error(w, "failed to decode request body", http.StatusBadRequest)
				return
			}
			newTodo.ID = len(todos) + 1
			todos = append(todos, newTodo)
			w.WriteHeader(http.StatusCreated)
		}
	}))

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
