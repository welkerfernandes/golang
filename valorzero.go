package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {

	//aqui são os valores zeros dos tipos variáveis sem atribuição
	// sendo que são variáveis package level scope
	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)
	fmt.Printf("%v , %T\n", c, c)
	fmt.Printf("%v , %T\n", d, d)

}
