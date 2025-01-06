package main

import (
	"fmt"
)

func create_pointer() *int {
	a := 1
	return &a
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

type Point struct {
	x float64
	y float64 
}

func main() {

	a := create_pointer()

	fmt.Printf("Hello, World! %d\n", a)

}
