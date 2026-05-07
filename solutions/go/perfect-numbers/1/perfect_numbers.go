package perfectnumbers

// Define the Classification type here.

import "errors"

type Classification string

const ClassificationDeficient Classification = "Deficient"
const ClassificationPerfect Classification = "Perfect"
const ClassificationAbundant Classification = "Abundant"

var ErrOnlyPositive = errors.New("Must be greater than 0")

func FindProperDivisors(n int64) []int {
	var divisorList []int
	for i := 1; i < int(n); i++ {
		if n % int64(i) == 0 {
			divisorList = append(divisorList, i)
		}
	}
	return divisorList
}

func SumOfProperDivisors(n int64) int {
	var sum int
	divisorList := FindProperDivisors(n)
	for i := 0; i < len(divisorList); i++ {
		sum += divisorList[i]
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	aliquotSum := SumOfProperDivisors(n)
	nInt := int(n)
	if n <= 0 {
		// FIXME: Is it ok to return a classification even on error ?
		return ClassificationPerfect, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}
		
	switch {
	case aliquotSum == nInt: return ClassificationPerfect, nil
	case aliquotSum < nInt: return ClassificationDeficient, nil
	case aliquotSum > nInt: return ClassificationAbundant, nil
	default: return ClassificationAbundant, ErrOnlyPositive
	}
}


