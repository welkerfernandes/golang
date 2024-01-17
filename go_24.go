package main

import (
	"fmt"
)

var x int

func main() {

	x = 100
	y := x << 2

	fmt.Printf(" decimal: %d \n binario: %b\n hexadecimal: %#X \n Tipo: %T\n\n", x, x, x, x)
	fmt.Printf(" decimal: %d \n binario: %b\n hexadecimal: %#X \n Tipo: %T", y, y, y, y)
}
