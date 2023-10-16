package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 16

	x := (a == b)
	fmt.Printf("%v == %v: %v\n", a, b, x)

	y := (a != b)
	fmt.Printf("%v != %v: %v\n", a, b, y)

	z := (a < b)
	fmt.Printf("%v < %v: %v\n", a, b, z)

	w := (a > b)
	fmt.Printf("%v > %v: %v", a, b, w)
}
