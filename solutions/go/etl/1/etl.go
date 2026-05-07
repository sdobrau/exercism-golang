package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	var outMap = map[string]int{}
	for point, slice := range in {
		for _, letter := range slice {
			lowerLetter := strings.ToLower(letter)
			outMap[lowerLetter] = point
		}
	}
	return outMap
}
