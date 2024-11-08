package main

import "fmt"

type ar2x2 [2][2]int

// Traditional Add() function
func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}

// Type method Add()
func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func main() {
	ax := ar2x2{[2]int{1, 2}, [2]int{3, 4}}
	ay := ar2x2{[2]int{5, 6}, [2]int{7, 8}}

	res := Add(ax, ay)
	fmt.Printf("res: %v\n", res)

	ax.Add(ay)
	fmt.Printf("ax: %v\n", ax)
}
