package diffsquares

func SquareOfSum(num int) int {
	tmp := (num + 1) * num / 2
	return tmp * tmp
}

func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
