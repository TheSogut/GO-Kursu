package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Computer", 500, "Lenovo"})
	fmt.Println(product{"Computer", 500, ""})
	fmt.Println(product{name: "Computer"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
