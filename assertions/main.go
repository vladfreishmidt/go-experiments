package main

import "fmt"

func main() {
	anInt := returnNumber()
	number, ok := anInt.(int)
	if ok {
		fmt.Println("Type assertion successful: ", number)
	} else {
		fmt.Println("Type assertion failed!")
	}
	number++
	fmt.Println(number)

	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	// _ = anInt.(bool) // panic - anInt is not bool
}

func returnNumber() interface{} {
	return 12
}
