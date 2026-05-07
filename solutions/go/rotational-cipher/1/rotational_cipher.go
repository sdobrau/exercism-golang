package rotationalcipher

import (
	"fmt"
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	alphabets := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}
	
	rotatedStr := ""
	
	for _, char := range plain {
		// for each char
		isDigit := unicode.IsDigit(char)
		isWhitespace := unicode.IsSpace(char)
		isPunctuation := unicode.IsPunct(char)

		fmt.Printf("current char is %c\n", char)
		if isDigit || isWhitespace || isPunctuation {
			rotatedStr += string(char)
		} else {
		rotationIndex := strings.Index(alphabets, string(char)) + shiftKey
		if rotationIndex != 0 { // if found 
			rotatedStr += string(alphabets[rotationIndex])
		} else { // insert itself
			rotatedStr += fmt.Sprintf("%s", string(char)) // whitespace or digit
		}
		}
	}
	return rotatedStr
}
