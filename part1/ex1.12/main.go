package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	imageURL       = "https://picsum.photos/1200"
	cacheFilePath  = "/cache/image.jpg"
	updateInterval = time.Hour
)

func fetchAndCacheImage() error {
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(cacheFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	if err := os.MkdirAll("/cache", os.ModePerm); err != nil {
		log.Fatalf("failed to create cache directory: %v", err)
	}

	if _, err := os.Stat(cacheFilePath); os.IsNotExist(err) {
		log.Println("no cached image found, fetching a new one...")
		if err := fetchAndCacheImage(); err != nil {
			log.Fatalf("failed to fetch initial image: %v", err)
		}
	}

	go func() {
		for range time.Tick(updateInterval) {
			log.Println("fetching a new image...")
			if err := fetchAndCacheImage(); err != nil {
				log.Printf("failed to fetch image: %v", err)
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, cacheFilePath)
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
