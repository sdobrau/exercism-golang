package squareroot

import (
	"errors"
)

func SquareRoot(number int) (int, error) {
	if number <= 0 {
		return 0, errors.New("0 is not a valid value")
	}
	i := 0
	for {
		if i * i == number {
			return i, nil
		}
		i++
	}
}
