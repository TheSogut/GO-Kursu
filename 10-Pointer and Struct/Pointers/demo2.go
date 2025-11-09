package pointers

import "fmt"

func Demo2(nums []int) {
	nums[0] = 100
	fmt.Println("The number in demo ", nums[0])
}
