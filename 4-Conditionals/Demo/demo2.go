package conditionals

import "fmt"

func Demo2() {
	var account float64 = 1000
	var wantToDraw float64 = 900

	if wantToDraw > account {
		fmt.Println("There is no enough money")
	} else if wantToDraw == account {
		fmt.Println("You're withdrawing all of your balance")
		account = account - wantToDraw
	} else {
		fmt.Println("Your money is preparing")
		account = account - wantToDraw
	}

	fmt.Println("End. Balance : " + fmt.Sprintf("%v", account))
}
