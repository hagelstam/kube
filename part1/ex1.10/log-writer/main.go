package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting home directory: %v", err)
	}

	logDir := filepath.Join(homeDir, "filelog")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	filePath := filepath.Join(logDir, "timestamp.txt")

	for {
		timestamp := fmt.Sprintf("%s: %d\n", time.Now().Format(time.RFC3339), rand.Int())

		err := os.WriteFile(filePath, []byte(timestamp), 0644)
		if err != nil {
			fmt.Printf("error writing file: %v", err)
		}

		time.Sleep(5 * time.Second)
	}
}
