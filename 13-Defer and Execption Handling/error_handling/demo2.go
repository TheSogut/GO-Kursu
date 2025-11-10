package errorhandling

import (
	"errors"
	"fmt"
)

func guessNumber(number int) (string, error) {
	target := 50
	if number < 1 || number > 100 {
		return "", errors.New("your number must be between 1 and 100")
	}
	if number > target {
		return "Your number is high", nil
	}
	if number < target {
		return "Your number is low", nil
	}
	return "You found it", nil
}

func Demo2() {
	fmt.Println(guessNumber(50))
	fmt.Println(guessNumber(101))
	fmt.Println(guessNumber(-10))

}
