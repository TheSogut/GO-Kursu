package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Product saved :", p.productName)
	defer Log()
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 5000}
	fmt.Println("Finished")
	fmt.Println(p.productName)
}

func Log() {
	fmt.Println("Logged")
}
