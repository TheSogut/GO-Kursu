package variables

import "fmt"

func Demo1() {
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
	fmt.Printf("data type : %T\n", tax3)

	// Logical Variables
	var status bool

	var txt1 string = "Ali"
	var txt2 string = "Ahmet"

	// status = txt1 == txt2
	status = txt1 != txt2

	fmt.Println(status)
	fmt.Println(!status)

}