package main

import (
	"fmt"
	"gocourse/channels"
)

func main() {
	// go goroutines.EvenNumbers()
	// go goroutines.OddNumbers()
	evenNumCn := make(chan int)
	oddNumCn := make(chan int)

	go channels.EvenNumbers(evenNumCn)
	go channels.OddNumbers(oddNumCn)

	evenNumSum, oddNumSum := <-evenNumCn, <-oddNumCn

	mul := evenNumSum * oddNumSum
	fmt.Println(mul)
}
