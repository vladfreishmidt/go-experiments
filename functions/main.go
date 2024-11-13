package main

import "fmt"

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

var (
	double int
	square int
)

func main() {
	var num int = 99
	double, square = doubleSquare(num)

	fmt.Printf("input: %d, double: %d, square: %d\n", num, double, square)
}
