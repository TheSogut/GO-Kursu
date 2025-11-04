package loop

import (
	"fmt"
	"math/rand"
)

// Guessing random number

func Demo3() {
	target := rand.Intn(100)
	estimated := 0
	for i := 0; i < 1; {
		fmt.Println("Guess a number")
		fmt.Scanln(&estimated)

		if estimated == target {
			fmt.Println("You found the number")
			i = 1
		} else if estimated < target {
			fmt.Println("Your guess is low")
		} else {
			fmt.Println("Your guess is high")
		}
	}

}
