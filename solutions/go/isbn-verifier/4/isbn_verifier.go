package isbnverifier

import (
	"strings"
	"unicode"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnLen := len(isbn)
	
	if isbnLen != 10 || isbn == "" { return false }
	
	sum := 0
	for i, char := range isbn {
		// if X and last 
		if string(char) == "X" && i == isbnLen - 1 { // last X is OK
			sum += 10
			
		} else if unicode.IsLetter(char) { // any other case is not accepted
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
