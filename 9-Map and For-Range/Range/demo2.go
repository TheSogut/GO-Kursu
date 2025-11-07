package for_range

import "fmt"

func Demo2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, num := range nums {
		if num%2 == 1 {
			sum += num
		}
	}
	fmt.Println(sum)
}
