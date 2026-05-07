package perfectnumbers

// Define the Classification type here.

import ( "errors" )

type Classification string

const (
	ClassificationDeficient  = "Deficient" 
	ClassificationPerfect  = "Perfect"
	ClassificationAbundant  = "Abundant"
	NoClassification  = "No classification"
)

var ErrOnlyPositive = errors.New("Must be greater than 0")

func SumOfProperDivisors(n int64) int64 {
	var sum int64
	// Let's find the proper divisors
	for i := int64(1); i < n; i++ {
		// If a divisor
		if n % i == 0 {
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


