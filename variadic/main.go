package main

import (
	"fmt"
	"os"
)

func main() {
	// put the arguments of a variadic function inline
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)

	// use a slice variable with the unpack operator
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)

	// convert []string into []interface{} in order to use the unpack operator
	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty...)

	// different way of converting []string into []interface{}
	arguments := os.Args[1:]
	empty = make([]interface{}, len(arguments))
	for i := range arguments {
		empty[i] = arguments[i]
	}
	everything(empty...)

	str := []string{"One", "Two", "Three"}
	everything(str, str, str)

}

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}

	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}
