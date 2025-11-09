package deferstatement

import "fmt"

func Demo2(num int32) string {
	defer DeferFunc()
	if num%2 == 0 {
		return "Even Number"
	}
	if num%2 != 0 {
		return "Odd Number"
	}

	return ""
}

func DeferFunc() {
	fmt.Println("Finished")
}

func Test() {
	result := Demo2(10)
	fmt.Println(result)
}
