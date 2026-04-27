package rnatranscription

import (
)



func ToRNA(dna string) string {
	dnaToRnaMap := map[string]rune{
		"G": 'C',
		"C": 'G',
		"T": 'A',
		"A": 'U',
	}
	runesSlice := []rune(dna)
	for i, v := range runesSlice {
		runesSlice[i] = rune(dnaToRnaMap[string(v)])
	}
	rnaString := string(runesSlice)
	return rnaString
}
