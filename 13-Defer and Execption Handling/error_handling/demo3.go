package errorhandling

import (
	"fmt"
)

type BorderException struct {
	parameter int
	message   string
}

func (b BorderException) Error() string {
	return fmt.Sprintf("%d-----%s", b.parameter, b.message)
}

func Guess2(number int) (string, error) {
	if number < 1 || number > 100 {
		return "", &BorderException{number, "Out of bounds"}
	}
	return "You found it", nil
}
