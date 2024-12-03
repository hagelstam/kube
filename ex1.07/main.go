package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	currentTime  string
	currentValue int
)

func updateLog() {
	for {
		currentValue = rand.Int()
		currentTime = time.Now().Format(time.RFC3339)
		fmt.Printf("%s: %d\n", currentTime, currentValue)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	go updateLog()

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s: %d", currentTime, currentValue)
	})

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
