package main

import "fmt"

// ISSUE #1: The function signature does not limit the allowed data types.
// ISSUE #2: Need to specifically handle each case, and writing more code.
// Compiler and developer do not have to guess many things with that code,
// which is not the case with generics, where the compiler and the runtime
// have more work to do. (This kind of work introduces delays in the execution time).
func Print(s interface{}) {
	switch s.(type) {
	case int:
		fmt.Println(s.(int) + 1)
	case float64:
		fmt.Println(s.(float64) + 1)
	default:
		fmt.Println("Unknown data type!")
	}
}

// PrintGenerics handles all available data types
func PrintGenerics[T any](s T) {
	fmt.Println(s)
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

// PrintNumeric function supports all numeric data types, with the use
// of Numeric constraint.
// No need to specifically add code for supporting each distinct data type.
func PrintNumeric[T Numeric](s T) {
	fmt.Println(s + 1)
}

func main() {
	Print(12)
	Print(-1.23)
	Print("Hi!")

	PrintGenerics(1)
	PrintGenerics("a")
	PrintGenerics(-2.33)

	PrintNumeric(1)
	PrintNumeric(-2.33)
}
