package conditionals

import "fmt"

func Demo3() {
	//Declare three integer variable
	//Print the largest one

	var number1, number2, number3 int = 50, 40, 30

	if number1 > number2 {
		if number1 > number3 {
			println("Largest number is " + fmt.Sprintf("%v", number1))
		} else {
			println("Largest number is " + fmt.Sprintf("%v", number3))
		}
	} else {
		if number2 > number3 {
			println("Largest number is " + fmt.Sprintf("%v", number2))
		} else {
			println("Largest number is " + fmt.Sprintf("%v", number3))
		}
	}
}
