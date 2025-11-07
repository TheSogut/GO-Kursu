package main

import (
	"fmt"
	functions "gocourse/Functions"
)

func main() {
	functions.SayHi("Enes")
	// functions.Sum(2, 3)
	var result = functions.Sum(2, 6)
	fmt.Println(result)
}
