package errorhandling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("error_handling/demo1.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File not found", pErr.Path)
			return
		} else {
			fmt.Println("File not found")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
