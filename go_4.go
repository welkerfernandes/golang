package main

import (
	"fmt"
)

// não pode mudat o tipo de uma variavel
var x int = 10

var y int

func main() {
	//aqui mudei o valor mas não mudei o tipo
	x = 20
	fmt.Printf("x = %v | %T\n\n\n", x, x)

	y = 10
	fmt.Printf("y = %v | %T", y, y)
}
