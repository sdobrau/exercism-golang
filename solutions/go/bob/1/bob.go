// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
	"strings"
	"regexp"
)

func IsQuestion(remark string) bool {
	// strip all whitespace first to check for ?
	pattern := "[\\s]" // perl whitespace
	re := regexp.MustCompile(pattern)
	remark = re.ReplaceAllString(remark, "")
	if remark[len(remark) - 1] == '?' {
		return true
	} else {
		return false
	}
}

func IsYelling(remark string) bool {
	if IsOnlyPunctOrDigit(remark) {
		return false
	} else {
		return remark == strings.ToUpper(remark) // if equal to upper
	}
}

func IsOnlyPunctOrDigit(remark string) bool {
	for _, char := range remark {
		if !unicode.IsPunct(char) && !unicode.IsDigit(char){
			return false
		}
	}
	return true
}

func IsSilence(remark string) bool {
	for _, char := range remark {
		// test for noise
		if  unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsPunct(char) {
			return false
		}
	}
	// if no noise found then silence
	return true
}

func NoLettersNoQuestion(remark string) bool {
	for _, char := range remark {
		if unicode.IsLetter(char) {
			return false
		}
	}
	if !IsQuestion(remark) {
		return false
	} else {
		return true
	}
}

func NoLettersQuestion(remark string) bool {
	for _, char := range remark {
		if unicode.IsLetter(char) {
			return false
		}
	}
	if IsQuestion(remark) {
		return true
	} else {
		return false
	}
}

func DigitsAndPunctOnly(remark string) bool {
	for _, char := range remark {
		if unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {

	switch {
	case IsSilence(remark):
		return "Fine. Be that way!"
	case NoLettersQuestion(remark):
		return "Sure."
	case DigitsAndPunctOnly(remark):
		return "Whatever."
	case NoLettersNoQuestion(remark):
		return "Whatever."
	case IsQuestion(remark) && !IsYelling(remark):
		return "Sure."
	case IsYelling(remark) && !IsQuestion(remark):
		return "Whoa, chill out!"
	case IsYelling(remark) && IsQuestion(remark) && !IsOnlyPunctOrDigit(remark):
		return "Calm down, I know what I'm doing!"
	default: return "Whatever."

	}
}
