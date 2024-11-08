package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

// Set values
type T struct {
	F1 int
	F2 string
	F3 float64
}

func main() {
	A := Record{"String Value", -12.123, Secret{"Vladyslav", "Freishmidt"}}
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())

	iType := r.Type()
	fmt.Printf("i Type %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// Convert to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// Using the internal value
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}

	// Set values
	fmt.Println("")

	A1 := T{1, "F2", 3.0}
	fmt.Println("A1:", A1)

	// Examine variable A1
	r1 := reflect.ValueOf(&A1).Elem()
	fmt.Println("String value:", r1.String())

	typeOfA1 := r1.Type()

	for i := 0; i < r.NumField(); i++ {
		f := r1.Field(i)
		tOfA1 := typeOfA1.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA1, f.Type(), f.Interface())

		k := reflect.TypeOf(r1.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r1.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r1.Field(i).SetString("Changed!")
		}
	}

	fmt.Println("A1:", A1)
}
