// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var included_person string
	if name == "" {
		included_person = "you"
	} else {
		included_person = name
	}
	return fmt.Sprintf("One for %s, one for me.", included_person)
}
