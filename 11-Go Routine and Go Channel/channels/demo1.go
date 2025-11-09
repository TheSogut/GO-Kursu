package channels

func EvenNumbers(evenNumCn chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum += i
	}

	evenNumCn <- sum
}
func OddNumbers(oddNumCn chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum += i
	}

	oddNumCn <- sum
}
