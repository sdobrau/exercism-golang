package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0
	for nr := 1; nr <= n; nr++ {
		sum += nr
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for nr := 1; nr <= n; nr++ {
		sum += nr * nr
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
