package isbnverifier

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")
	fmt.Sprintf("isbn is %s and of length %d\n", isbn, len(isbn))
	if len(isbn) != 10 || isbn == "" { return false } 
	sum := 0
	
	match, _ := regexp.MatchString("[ABCDEFGHIJKLMNOPQRSTUVWYZ]", isbn)
	if match {
		return false
	}
	
	for i:= 0; i < len(isbn); i++ {
		// if X and last 
		if string(isbn[i]) == "X" && i == len(isbn) - 1 {
			sum += 10
		}
		
		digit, err := strconv.Atoi(string(isbn[i]))
		
		if err != nil {
			// the conversion did not work so it is a non-digit character
			sum += 0
		} else {
			fmt.Printf("multiplying %d by %d\n", digit, (len(isbn) - i))
			sum += int(digit) * (len(isbn) - i)
			fmt.Printf("sum is %d\n", sum)
		}
	}
	
	fmt.Printf("Sum of isbn %s is %d\n", isbn, sum)
	if sum % 11 == 0 {
		fmt.Printf("ISBN %s is valid\n", isbn)
		return true
	} else {
		fmt.Printf("ISBN %s is not valid\n", isbn)
		return false
	}
}
