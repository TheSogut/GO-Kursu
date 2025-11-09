package pointers

import "fmt"

func Demo1(num *int) {
	*num = *num + 1
	fmt.Println("The number in demo ", num)
}
