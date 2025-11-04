package loop

import (
	"fmt"
	"math/rand"
)

// Guessing random number

func Demo3() {
	target := rand.Intn(100)
	guess := 0
	guessCount := 0
	for i := 0; i < 1; {
		fmt.Println("Guess a number")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("You found the number")
			i = 1
		} else if guess < target {
			fmt.Println("Your guess is low")
		} else {
			fmt.Println("Your guess is high")
		}
		guessCount++
	}

	if guessCount <= 3 {
		fmt.Println("Excellent")
	} else if guessCount <= 6 {
		fmt.Println("Good")
	} else {
		fmt.Println("Not Bad")
	}
	fmt.Println("Guess Count : " + fmt.Sprintf("%v", guessCount))
}
