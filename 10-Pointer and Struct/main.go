package main

import (
	"fmt"
	pointers "gocourse/Pointers"
)

func main() {
	// num := 20
	// pointers.Demo1(&num)
	// fmt.Println("The number in main ", num)

	nums := []int{1, 2, 3}
	pointers.Demo2(nums)
	fmt.Println("The number in main ", nums[0])
}
