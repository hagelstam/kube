package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		fmt.Printf("%s: %d\n", time.Now().Format(time.RFC3339), rand.Int())
		time.Sleep(5 * time.Second)
	}
}
