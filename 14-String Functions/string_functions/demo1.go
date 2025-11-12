package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	name := "Enes"
	fmt.Println(s.Count(name, "e"))
	fmt.Println(s.Contains(name, "n"))
	fmt.Println(s.Index(name, "n"))
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))
}
