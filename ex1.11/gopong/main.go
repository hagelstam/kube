package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting home directory: %v", err)
	}

	logDir := filepath.Join(homeDir, "files")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	filePath := filepath.Join(logDir, "pong.txt")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, _ := os.ReadFile(filePath)
		count, _ := strconv.Atoi(string(data))
		count++

		os.WriteFile(filePath, []byte(strconv.Itoa(count)), 0644)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pong %d", count-1)))
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
