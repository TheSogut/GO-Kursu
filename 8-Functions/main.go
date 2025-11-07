package main

import (
	"fmt"
	functions "gocourse/Functions"
)

func main() {
	functions.SayHi("Enes")
	// functions.Sum(2, 3)
	// var result = functions.Sum(2, 6)
	// fmt.Println(result)
	// result1, result2, result3, result4 := functions.FourOperations(10, 6)
	// fmt.Println("Sum :", result1)
	// fmt.Println("Sub :", result2)
	// fmt.Println("Mul :", result3)
	// fmt.Println("Div :", result4)

	fmt.Println(functions.SumVariadic(1, 4, 6, 3, 10))
	fmt.Println(functions.SumVariadic(1, 4))
	fmt.Println(functions.SumVariadic())
}
