package loop

import "fmt"

func Demo1() {
	var text string = "Hello World!"

	i := 1

	for i <= 5 {
		fmt.Println(text)
		i = i + 1
	}

	fmt.Println("End")

}
