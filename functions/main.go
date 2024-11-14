package main

import "fmt"

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	n := 10
	d, s := doubleSquare(n)
	fmt.Printf("input: %d, double: %d, square: %d\n", n, d, s)

	// An anonymous function
	anF := func(param int) int {
		return param * param
	}
	fmt.Println("anF of", n, "is", anF(n))
}
