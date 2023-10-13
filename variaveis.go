package main

import (
	"fmt"
)

func main() {
	//operado curto declaração

	x := 16
	y := "string"

	// o %v mostra o valor e o %T mostra o tipo de dado sendo que estamos utilizando o print format ou (printf)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	//operador de atribuição =  eo operador de declaração :=
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)
	// se eu tiver duas variáveis sendo que uma tem uma declaração (variável nova com valor) ai pode usar o gopher (:=) em vez do (=)
	x, z := 30, 59
	fmt.Printf("x %v, %T\n", x, x)
	fmt.Printf("x %v, %T", z, z)
}
