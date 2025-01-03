package main

import (
	"fmt"

	"github.com/csci780example/GoTutorial/math/simplemath"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	a:= 1
	b := 2
	c := simplemath.Max(a, b)

	fmt.Printf("Hello, World! %d\n", c)
}