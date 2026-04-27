package resistorcolor

import "fmt"


var colorMap = map[string]int{
	"black": 0,
	"brown": 1,
	"red": 2,
	"orange": 3,
	"yellow": 4,
	"green": 5,
	"blue": 6,
	"violet": 7,
	"grey": 8,
	"white": 9,
}

// Colors returns the list of all colors.
func Colors() []string {
	var outputString []string
	for k, _ := range colorMap {
		outputString = append(outputString, fmt.Sprintf("%s", k))
	}
	return outputString
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colorMap[color]
}
