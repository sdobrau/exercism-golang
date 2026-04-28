package eliudseggs

import (
	"fmt"
)

func EggCount(displayValue int) int {
	//print number in base2 using printf
	eggs := 0
	binRepr := fmt.Sprintf("%b", displayValue)
	for _, v := range binRepr {
		if v == '1' {
			eggs += 1
		}
	}
	return eggs
}
