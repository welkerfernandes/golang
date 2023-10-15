package main

import (
	"fmt"
)

var x bool

func main() {

	fmt.Println("Valor Zero: ", x) // value zero
	x = true
	fmt.Println("Valor Atribuido por mim: ", x) // valor atribuido
	x = 10 > 100
	y := 10 == 100
	z := 10 < 100

	fmt.Println("10 é maior que 100?", x, "\n")
	fmt.Println("10 é igual a 100?", y, "\n")
	fmt.Println("10 é menor que 100?", z)

}
