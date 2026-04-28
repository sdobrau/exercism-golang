package isbnverifier

import (
	"strings"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnLen := len(isbn)
	lastCharIndex := isbnLen - 1
	if isbnLen != 10 {
		return false
	}
	sum := 0
	for i, char := range isbn {
		// if X and last 
		if char == 'X' && i == lastCharIndex { // last X is OK
			sum += 10
		} else if currentChar, isALetter := strconv.Atoi(string(char)); isALetter != nil { // any other case is not accepted
			return false
		} else {
			sum += currentChar * (isbnLen - i)
		}
	}	
	return sum % 11 == 0 
}
