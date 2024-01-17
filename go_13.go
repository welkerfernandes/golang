// fundamentos da programção

package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println("X está sem valor ", x) //zero value
	x = true
	fmt.Println("X foi atribuido um valor ", x) //valor atribuido
	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)

	fmt.Println("\n", x)
	fmt.Println("\n", y)
	fmt.Println("\n", z)
}
