package armstrongnumbers

import (
	"strconv"
	"fmt"
	"math"
)

func IsNumber(n int) bool {
	sum := 0.0
	toThePowerOf := float64(len(strconv.Itoa(n)))
	for i := 0; i < len(strconv.Itoa(n)); i++ {
		nthDigit := strconv.Itoa(n)[i] - '0'
		sum += math.Pow(float64(nthDigit), toThePowerOf)
	}
	if sum == float64(n) {
		return true
	} else {
		return false
	}
}

func main() {
	IsNumber(153)
}
