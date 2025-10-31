package main

import "fmt"

func main() {

	// String variable
	var txt string = "Hello World!"
	fmt.Print(txt)
	fmt.Print(txt)
	fmt.Print(txt)
	fmt.Print(txt)
	fmt.Println(txt)

	// integer variable
	var tax int = 15
	fmt.Println(tax)
	fmt.Println(100 + (100 * tax / 100))

	// Float Variable
	var tax2 float32 = 0.1
	fmt.Println(tax2)
	fmt.Println(100 + (100 * int(tax2)))

	// Another Variaton of Variable Defination
	tax3 := 25.2
	fmt.Println(tax3)
	fmt.Printf("data type : %T", tax3)

}
