package isbnverifier

import (
	"strconv"
	"strings"
	"fmt"
)

func IsValidISBN(isbn string) bool {
	// this code works, however the tests fail:
	//--- FAIL: TestIsValidISBN (0.00s)
	//--- FAIL: TestIsValidISBN/valid_isbn (0.00s)
        //isbn_verifier_test.go:12: IsValidISBN("3-598-21508-8")=false, want: true
	//--- FAIL: TestIsValidISBN/valid_isbn_with_a_check_digit_of_10 (0.00s)
	// isbn_verifier_test.go:12: IsValidISBN("3-598-21507-X")=false, want: true
	// both values return true
	strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 || isbn == "" { return false }
	sum := 0
	for i:= 0; i < len(isbn); i++ {
		if string(isbn[i]) == "X" {
			sum += 10 
		} else {		
			digit, _ := strconv.Atoi(string(isbn[i]))
			fmt.Printf("multiplying %d by %d\n", digit, (len(isbn) - i))
			sum += digit * (len(isbn) - i)
			fmt.Printf("sum is %d\n", sum)
		}
	}
	fmt.Printf("Sum is now before 11ing %d\n", sum)
	if sum % 11 == 0 {
		fmt.Printf("ISBN %s is valid\n", isbn)
		return true
	} else {
		fmt.Printf("ISBN %s is not valid\n", isbn)
		return false
	}
}
