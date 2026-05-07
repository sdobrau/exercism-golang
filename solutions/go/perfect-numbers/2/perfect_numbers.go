package perfectnumbers

// Define the Classification type here.

import (
	"errors"
)

type Classification string

const (
	// Should these be untyped?
	ClassificationDeficient Classification = "Deficient" 
	ClassificationPerfect Classification = "Perfect"
	ClassificationAbundant Classification = "Abundant"
	NoClassification Classification = "No classification"
)

var ErrOnlyPositive = errors.New("Must be greater than 0")

func FindProperDivisors(n int64) []int64 {
	var divisors []int64
	for i := int64(1); i < n; i++ { 
		if n % i == 0 {
		divisors = append(divisors, int64(i))
	}
	}
	return divisors
}

func SumOfProperDivisors(n int64) int64 {
	var sum int64
	divisors := FindProperDivisors(n)
	for _, v := range divisors {
		sum += v
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
	}
	// Is it idiomatic to catch anything else this way?
	return NoClassification, nil
}


