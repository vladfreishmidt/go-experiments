package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}
type S2 struct {
	F3 float64
}

func Print(s interface{}) {
	fmt.Println(s)
}

func main() {
	s1 := S1{12, "foo"}
	s2 := S2{3.12}

	Print(s1)
	Print(s2)
	Print("bar")
	Print(123)
}
