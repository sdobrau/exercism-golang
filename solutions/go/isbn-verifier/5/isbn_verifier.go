package isbnverifier

import (
	"strings"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnLen := len(isbn)
	lastCharIndex := len(isbn) - 1
	
	if isbnLen != 10 {
		return false
	}
	
	sum := 0
	for i, char := range isbn {
		// test if it's a letter using the error from strconv.Atoi
		_, IsLetterError := strconv.Atoi(string(char))
		
		// if X and last 
		if char == 'X' && i == lastCharIndex { // last X is OK
			sum += 10
			
		} else if IsLetterError != nil { // any other case is not accepted
			return false
			
		} else {
			digit, _ := strconv.Atoi(string(char))
			sum += digit * (isbnLen - i)
		}
	}
	
	if sum % 11 == 0 {
		return true
	} else {
		return false
	}
}
