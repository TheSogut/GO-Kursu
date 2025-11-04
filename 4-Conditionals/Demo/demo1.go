package conditionals

import "fmt"

func Demo1() {
	var account float64 = 1000
	var wantToDraw float64 = 900

	if wantToDraw > account {
		fmt.Println("There is no enough money")
	}

	if wantToDraw <= account {
		fmt.Println("Your money is preparing")
		account = account - wantToDraw
	}
	fmt.Println("End. Balance : " + fmt.Sprintf("%v", account))
}