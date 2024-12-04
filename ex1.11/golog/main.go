package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	currentTime string
	currentHash int
)

func updateLog() {
	for {
		currentHash = rand.Int()
		currentTime = time.Now().Format(time.RFC3339)
		fmt.Printf("%s: %d\n", currentTime, currentHash)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	go updateLog()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting home directory: %v", err)
	}

	counterFile := filepath.Join(homeDir, "files", "pong.txt")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := os.ReadFile(counterFile)
		count, _ := strconv.Atoi(string(data))

		fmt.Fprintf(w, "%s: %d\nPing / Pongs: %d", currentTime, currentHash, count)
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
