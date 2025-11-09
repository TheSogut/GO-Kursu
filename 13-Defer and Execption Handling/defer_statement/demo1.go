package deferstatement

import "fmt"

func A() {
	fmt.Println("a function has worked")
}
func C() {
	fmt.Println("c function has worked")
}
func D() {
	fmt.Println("d function has worked")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b function has worked")
}
