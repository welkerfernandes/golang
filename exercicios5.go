package main

import (
	"fmt"
)

type welker int

var x welker
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("Valor Atual: %v | Tipo: %T | Valor atribuido:", x, x)
	x := 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v | %T", y, y)
}
