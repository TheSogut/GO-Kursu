package interfaces

import "fmt"

func submit(i interface{}) {
	num, ok := i.(int)
	fmt.Println(num, ok)
}

func Demo3() {
	var num1 interface{} = "Enes"
	submit(num1)

	var num2 interface{} = 5
	submit(num2)
}
