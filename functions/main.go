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
	fmt.Println(sortTwo(1, -3))
	fmt.Println(sortTwo(-1, 0))

}

func sortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

// Named return values
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}

	min = x
	max = y
	return
}
