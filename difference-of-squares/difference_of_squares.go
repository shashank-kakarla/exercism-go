// Package diffsquares has functions which calculate the square of integers,
// sum of the squares of integers and square of the sum of integers
package diffsquares

// CalculateSquare calculates the square of the given integer.
func CalculateSquare(number int) int {
	return number * number
}

// SquareOfSum calculates the square of the sum of integers from 1 to N.
func SquareOfSum(N int) int {
	sum := (N * (N + 1)) / 2
	return CalculateSquare(sum)
}

// SumOfSquares calculates the sum of squares of integers from 1 to N.
func SumOfSquares(N int) int {
	sum := (N * (N + 1) * (2*N + 1)) / 6
	return sum
}

// Difference returns the difference between the values returned
// by SquareOfSum and SumOfSquares respectively.
func Difference(N int) int {
	return SquareOfSum(N) - SumOfSquares(N)
}
