package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
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

func getPong() (string, error) {
	resp, err := http.Get("http://gopong-svc:3001/pingpong")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	counterBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(counterBytes), nil
}

func main() {
	go updateLog()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pong, err := getPong()
		if err != nil {
			http.Error(w, "failed to fetch pong", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s: %d\n%s", currentTime, currentHash, pong)
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
