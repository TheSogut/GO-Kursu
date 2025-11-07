package functions

func FourOperations(num1 int, num2 int) (int, int, int, float32) {
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := float32(num1 / num2)

	return sum, sub, mul, div
}
