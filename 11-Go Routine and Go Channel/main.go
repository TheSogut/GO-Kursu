package main

import (
	"gocourse/goroutines"
	"time"
)

func main() {
	go goroutines.EvenNumbers()
	go goroutines.OddNumbers()
	time.Sleep(5 * time.Second)
}
