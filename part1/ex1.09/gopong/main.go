package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable required but not set")
	}

	var counter uint32 = 0

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		count := atomic.AddUint32(&counter, 1)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("pong %d", count-1)))
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
