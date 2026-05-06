package perfectnumbers

// Define the Classification type here.

import "errors"

type Classification int

const (
	ClassificationDeficient = iota // 0
	ClassificationPerfect          // 1
	ClassificationAbundant         // 2
	NoClassification               // 3
)

var ErrOnlyPositive = errors.New("Must be greater than 0")

func SumOfProperDivisors(n int64) int64 {
	var sum int64
	// Let's find the proper divisors
	for i := int64(1); i < n; i++ {
		// If a divisor
		if n%i == 0 {
			sum += i // Add it to the sum
		}
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return NoClassification, ErrOnlyPositive
	}
	aliquotSum := SumOfProperDivisors(n)
	switch {
	case aliquotSum == n:
		return ClassificationPerfect, nil
	case aliquotSum < n:
		return ClassificationDeficient, nil
	case aliquotSum > n:
		return ClassificationAbundant, nil
	default:
		return NoClassification, errors.New("Invalid value")
	}
}
