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
		// test if it's a letter using the error from strconv.Atoi
		stringChar, isLetterError := strconv.Atoi(string(char))
		
		// if X and last 
		if char == 'X' && i == lastCharIndex { // last X is OK
			sum += 10
			
		} else if isLetterError != nil { // any other case is not accepted
			return false
			
		} else {
			digit := stringChar
			sum += digit * (isbnLen - i)
		}
	}
	
	if sum % 11 == 0 {
		return true
	} else {
		return false
	}
}
