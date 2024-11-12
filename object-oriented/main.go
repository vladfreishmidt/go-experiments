package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

// Structure compose gets the fields of structure a
type compose struct {
	field1 int
	a
}

// Different structures can have methods with the same name
func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC c = c{a{120, 12}, b{"-12", -12}} // iC.A.A(), iC.B.A()

	iComp := compose{123, a{456, 789}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1) // Access the anonymous structure directly

	iC.bar()
	processA(iC)
}
