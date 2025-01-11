package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	port := 3001
	var counter uint32 = 0

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		count := atomic.AddUint32(&counter, 1)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Ping / Pongs: %d", count-1)))
	})

	fmt.Printf("server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
