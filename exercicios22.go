package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 16

	x := (a == b)
	fmt.Println(x)

	y := (a != b)
	fmt.Println(y)

	z := (a < b)
	fmt.Println(z)

	w := (a > b)
	fmt.Println(w)
}
