package atbashcipher

import (
	"strings"
	"slices"
	"unicode"
)

var atbashMap = map[rune]rune{
	'A': 'Z', 'B': 'Y', 'C': 'X', 'D': 'W', 'E': 'V',
	'F': 'U', 'G': 'T', 'H': 'S', 'I': 'R', 'J': 'Q',
	'K': 'P', 'L': 'O', 'M': 'N', 'N': 'M', 'O': 'L',
	'P': 'K', 'Q': 'J', 'R': 'I', 'S': 'H', 'T': 'G',
	'U': 'F', 'V': 'E', 'W': 'D', 'X': 'C', 'Y': 'B',
	'Z': 'A',
	'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v',
	'f': 'u', 'g': 't', 'h': 's', 'i': 'r', 'j': 'q',
	'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l',
	'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g',
	'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b',
	'z': 'a',
}

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

func Atbash(s string) string {
	var pt string
	var ct strings.Builder
	pt = strings.ReplaceAll(s, " ", "")
	for _, char := range pt {
		if unicode.IsDigit(char) {
			ct.WriteString(string(char))
		} else if unicode.IsPunct(char) {
			ct .WriteString("")
		} else {
			ct .WriteString(string(atbashMap[char]))
		}
	}
	return SplitStringIntoFivedChunks(strings.ToLower(ct.String()))
}


