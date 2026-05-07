// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
)


func getFirstLetter(s string) string {
	upperLetter := strings.ToUpper(string(s[0]))
	return upperLetter
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// first replace consecutive delim " - " with "-"
	fmt.Println("hello")
	hyphened := strings.ReplaceAll(s, " - ", "-")
	hyphened = strings.ReplaceAll(hyphened, " ", "-")
	// test cases
	// replace punctuation
	hyphened = strings.ReplaceAll(hyphened, ", ", "-")
	// remove apostrophes
	hyphened = strings.ReplaceAll(hyphened, "'", "")
	// Remove underscores
	hyphened = strings.ReplaceAll(hyphened, "_", "")
	fmt.Sprintf("%s\n", s)
	sliced := strings.Split(hyphened, "-")
	acronymString := ""
	for i, _ := range sliced {
		acronymString += getFirstLetter(sliced[i])
	}
	fmt.Printf("%s\n", acronymString)
	return acronymString
}
