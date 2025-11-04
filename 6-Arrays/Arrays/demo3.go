package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numbers)
	largest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}
	fmt.Println(largest)
}
