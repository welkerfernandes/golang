package main

import (
	"fmt"
)

func main() {

	x := 16

	fmt.Printf("x: %v,%T", x, x)

	x, z := 20, 30
	fmt.Printf("x: %v,%T", x, x)
	fmt.Printf("z: %v,%T", z, z)

}
