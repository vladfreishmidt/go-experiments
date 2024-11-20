package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	fmt.Println("************ doubleSquare() function call ************")
	// doubleSquare returns two int values
	n := 10
	d, s := doubleSquare(n)
	fmt.Printf("input: %d, double: %d, square: %d\n", n, d, s)

	fmt.Println("************ anF() function call ************")
	// An anonymous function
	anF := func(param int) int {
		return param * param
	}
	fmt.Println("anF of", n, "is", anF(n))
	fmt.Println(sortTwo(1, -3))
	fmt.Println(sortTwo(-1, 0))

	data := []Grades{
		{"John", "Doe", 10},
		{"Sarah", "Smith", 3},
		{"Mike", "Anderson", 8},
	}

	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("It is NOT sorted!")
	}
	sort.Slice(data,
		func(i, j int) bool { return data[i].Grade < data[j].Grade })
	fmt.Println("By Grade:", data)

	fmt.Println("************ funRet() function call ************")
	// Calling funRet
	n = 10
	i := funRet(n)
	j := funRet(-4)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))

	fmt.Println("************ d1() function call ************")
	d1()
}

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// sortTwo returns two int values sorting from smaller to bigger value
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

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}

	return func(k int) int {
		return k * k
	}
}

func d1() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}

	fmt.Println("Some d1 functionality")
}
