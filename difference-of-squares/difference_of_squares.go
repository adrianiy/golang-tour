// Package diffsquares contains methods to check difference of squares.
package diffsquares

// SquareOfSum sum all first n numbers and squares them.
func SquareOfSum(n int) (square int) {
	square = (n * (n + 1)) / 2
	square *= square

	return
}

// SumOfSquares returns the sum every first n numbers.
func SumOfSquares(n int) (sum int) {
	sum = (n * (n + 1) * (2*n + 1)) / 6

	return
}

// Difference return the difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
