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
		time.Sleep(5 * time.Second)
	}
}

func getPong() (string, error) {
	resp, err := http.Get("http://gopong-svc.dwk.svc.cluster.local:80/pingpong")
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

func getFileContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	go updateLog()

	port := 3000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pong, err := getPong()
		if err != nil {
			http.Error(w, "failed to fetch pong", http.StatusInternalServerError)
			return
		}

		text, err := getFileContent("/etc/config/information.txt")
		if err != nil {
			http.Error(w, "failed to fetch pong", http.StatusInternalServerError)
			return
		}

		file := fmt.Sprintf("file content: %s\n", text)
		env := fmt.Sprintf("env variable: MESSAGE=%s\n", os.Getenv("MESSAGE"))
		timestamp := fmt.Sprintf("%s: %d\n", currentTime, currentHash)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s%s%s%s", file, env, timestamp, pong)
	})

	fmt.Printf("server running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
