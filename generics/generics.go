package main

import "fmt"

// PrintSlice prints each element in the provided slice to the standard output.
// It accepts a slice of any data type using Go's generic type parameter T.
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	xi := []int{64, 128, 256, 512}
	xs := []string{"Server uptime: 152 days, 6 hours, 42 minutes.", "Database connection established on port 5432."}

	PrintSlice(xi)
	PrintSlice(xs)
}
