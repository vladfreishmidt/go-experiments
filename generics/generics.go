package main

import (
	"fmt"

	"github.com/vladfreishmidt/go-experiments/generics/constraints"
	genericsingostructures "github.com/vladfreishmidt/go-experiments/generics/generics-in-go-structures"
)

// PrintSlice prints each element in the provided slice to the standard output.
// It accepts a slice of any data type using Go's generic type parameter T.
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Same returns true if the two provided values are equal, and false otherwise.
// The function accepts any comparable type, allowing comparisons between values
// of the same type, such as integers, strings, or custom comparable types.
func Same[T comparable](a, b T) bool {
	return a == b
}

func main() {
	xi := []int{64, 128, 256, 512}
	xs := []string{"Server uptime: 152 days, 6 hours, 42 minutes.", "Database connection established on port 5432."}

	PrintSlice(xi)
	PrintSlice(xs)

	fmt.Println(Same([3]int{1, 2, 99}, [3]int{1, 2, 3})) // You are allowed to compare two arrays

	type aInt int
	type bInt int

	var xx aInt = 123
	// var yy bInt = 777

	sum := constraints.AddElements([]aInt{xx, xx})
	fmt.Printf("sum: %v\n", sum)

	var myList genericsingostructures.List[int]
	fmt.Println(myList)

	myList.Add(12)
	myList.Add(9)
	myList.Add(3)
	myList.Add(9)

	cur := myList.Start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}

		cur = cur.Next
	}

}
