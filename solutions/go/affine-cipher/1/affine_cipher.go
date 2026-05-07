package affinecipher

import (
	"errors"
	"slices"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var letterMap = map[rune]int{
	'a': 0,
	'A': 0,
	'b': 1,
	'B': 1,
	'c': 2,
	'C': 2,
	'd': 3,
	'D': 3,
	'e': 4,
	'E': 4,
	'f': 5,
	'F': 5,
	'g': 6,
	'G': 6,
	'h': 7,
	'H': 7,
	'i': 8,
	'I': 8,
	'j': 9,
	'J': 9,
	'k': 10,
	'K': 10,
	'l': 11,
	'L': 11,
	'm': 12,
	'M': 12,
	'n': 13,
	'N': 13,
	'o': 14,
	'O': 14,
	'p': 15,
	'P': 15,
	'q': 16,
	'Q': 16,
	'r': 17,
	'R': 17,
	's': 18,
	'S': 18,
	't': 19,
	'T': 19,
	'u': 20,
	'U': 20,
	'v': 21,
	'V': 21,
	'w': 22,
	'W': 22,
	'x': 23,
	'X': 23,
	'y': 24,
	'Y': 24,
	'z': 25,
	'Z': 25,
}


var codeMap = map[int]rune{
        0: 'a',
        1: 'b',
        2: 'c',
        3: 'd',
        4: 'e',
        5: 'f',
        6: 'g',
        7: 'h',
        8: 'i',
        9: 'j',
        10: 'k',
        11: 'l',
        12: 'm',
        13: 'n',
        14: 'o',
        15: 'p',
        16: 'q',
        17: 'r',
        18: 's',
        19: 't',
        20: 'u',
        21: 'v',
        22: 'w',
        23: 'x',
        24: 'y',
        25: 'z',
}

// Given a string, it splits it at each 5 character boundary so that
// "abcdefghijk" → "abcde fghij k"
func SplitStringIntoFivedChunks(text string) string {
	if len(text) > 5 {
		textSlice := []rune(text)
		for i := 0; i < len(textSlice); i += 6 {
			textSlice = slices.Insert(textSlice, i, ' ')
		}
		textSlice = slices.Delete(textSlice, 0, 1) // pop
		textString := string(textSlice)
		return textString
	} else {
		return text
	}
}

// Find the GCD of two sorted slices of divisors.
func FindGreatestCommonDivisor(a, b []int) int {
	// the divisors are already sorted by the FindDivisors
	// operation, no sorting needed
	var commonDivisors []int
	for _, v := range a {
		for _, x := range b {
			if v == x {
				commonDivisors = append(commonDivisors, v)
			}
		}
	}
	return commonDivisors[len(commonDivisors)-1]
}

// Return the divisors of int n, including n itself
func FindDivisors(n int64) []int {
	var divisorList []int
	for i := 1; i < int(n); i++ {
		if n%int64(i) == 0 {
			divisorList = append(divisorList, i)
		}
	}
	divisorList = append(divisorList, int(n))
	return divisorList
}

// Given two numbers, return whether they are coprime. For two
// numbers to be coprime their GCD (Greatest common divisor) must be
// 1.
func IsCoprime(a, b int) bool {
	var firstDivisors = FindDivisors(int64(a))
	var secondDivisors = FindDivisors(int64(b))
	if FindGreatestCommonDivisor(firstDivisors, secondDivisors) == 1 {
		fmt.Println("Is coprime")
		return true
	}
	fmt.Println("Is not coprime")
	return false

}

var ErrNotCoprime = errors.New("Are not coprime")

// 
func Encode(text string, a, b int) (string, error) {
	if !IsCoprime(a, 26) {
		return "", ErrNotCoprime
	}
	var cipherText strings.Builder
	text = strings.ReplaceAll(text," ", "") // trim whitespace

	// trim punctuation
	pattern := "[\\p{P}]"
	re := regexp.MustCompile(pattern)
	text = re.ReplaceAllString(text, "")
	
	for _, char := range text {
		_, isNotDigit := strconv.Atoi(string(char))
		if isNotDigit != nil { //if letter
			encValue := (a * letterMap[char] + b) % 26
			cipherText.WriteString(string(codeMap[encValue]))
		} else if isNotDigit == nil { // is digit, just append
			cipherText.WriteString(string(char))
		}		
	}
	
	return SplitStringIntoFivedChunks(cipherText.String()), nil
}

func FindMMIOfA(a int) int {
	i := 0
	for {
		if (15*i)%26 == 1 {
			return i
		} else {
			i++
		}
	}
}

func Decode(text string, a, b int) (string, error) {
	var plainText strings.Builder
	if!IsCoprime(a, 26) {
		return "", errors.New("The numbers are not coprime. Abandoning")
	}
	text = strings.ReplaceAll(text," ", "") // trim whitespace
	
	for _, char := range text {
		// D(y) = (a^-1)(y - b) mod m
		encryptedLetterValue := letterMap[char] // y 
		fmt.Printf("The value of the encrypted letter %c is %d\n", char, encryptedLetterValue)

		// a^-1 = FindMMIOfA(a)
		// TOFIX: what is wrong with this formula? it returns negative values
		decValue := (FindMMIOfA(a) * (encryptedLetterValue - b)) % 26 
		fmt.Printf("Multiplying %d by %d then mod 26\n", FindMMIOfA(a), encryptedLetterValue-b)
		fmt.Printf("%d mod 26\n", FindMMIOfA(a)*(encryptedLetterValue-b))
		fmt.Printf("decValue is %v\n", decValue)
		fmt.Printf("Dec letter is %s\n", string(codeMap[decValue]))
		plainText.WriteString(string(codeMap[decValue]))
		fmt.Printf("plaintext is %s\n", plainText.String())
	}
	return plainText.String(), nil
}
