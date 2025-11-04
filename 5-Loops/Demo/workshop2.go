package loop

import "fmt"

func Demo4() {
	// Take a number from user
	// Check is it prime or not

	uNumber := 0
	isPrime := false
	fmt.Println("Enter a number")
	fmt.Scanln(&uNumber)

	for i := 2; i < uNumber/2; i++ {
		if uNumber%i == 0 {
			isPrime = true
		}
	}
	if isPrime {
		fmt.Println("Your number is not a prime number")
	} else {
		fmt.Println("Your number is a prime number")
	}
}
