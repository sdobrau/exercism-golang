package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, "-","")
	word = strings.ReplaceAll(word, " ","")
	for i, _ := range word {
		for j := i; j < len(word) - 1; j++ {
			if strings.EqualFold(string(word[j+1]), string(word[i])) {
				return false
			}
		}
	}
	return true 
}
