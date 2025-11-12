package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Enes"
	fmt.Println(s.HasPrefix(name, "En"))
	fmt.Println(s.HasSuffix(name, "En"))
	fmt.Println(s.Index(name, "En"))
	letters := []string{"e", "n", "e", "s"}
	fmt.Println(s.Join(letters, "*"))

	fmt.Println(s.Replace(name, "En", "EN", -1))
	fmt.Println(s.Split(name, "e"))
	fmt.Println(s.Repeat(name, 3))
}
