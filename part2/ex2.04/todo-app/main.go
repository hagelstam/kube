package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	imageURL       = "https://picsum.photos/1200"
	cacheFilePath  = "/cache/image.jpg"
	updateInterval = time.Hour
	todoAPI        = "http://todo-backend-svc:3001/todos"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

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

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, cacheFilePath)
	})

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(todoAPI)
		if err != nil {
			log.Printf("error fetching todos: %v", err)
			tmpl.Execute(w, nil)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error reading response: %v", err)
			tmpl.Execute(w, nil)
			return
		}

		var todosResp struct {
			Todos []Todo `json:"todos"`
		}
		err = json.Unmarshal(body, &todosResp)
		if err != nil {
			log.Printf("error parsing todos: %v", err)
			tmpl.Execute(w, nil)
			return
		}

		sort.Slice(todosResp.Todos, func(i, j int) bool {
			return todosResp.Todos[i].ID > todosResp.Todos[j].ID
		})

		tmpl.Execute(w, todosResp.Todos)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		todoText := r.FormValue("todo")
		if strings.TrimSpace(todoText) == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		todo := Todo{Text: todoText}
		jsonData, _ := json.Marshal(todo)

		resp, err := http.Post(todoAPI, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("error creating todo: %v", err)
		}
		defer resp.Body.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Printf("server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
