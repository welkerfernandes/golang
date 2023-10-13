package main

import (
	"fmt"
)

// aqui criei meu proprio tipo e chamei ele de hotdog
type hotdog int

// aqui atribui um valor para meu tipo e utilizei ele em uma variável
var b hotdog = 10

func main() {
	// aqui atribui um valor int para uma variável x
	x := 10

	//aqui pedi para mostrar o valor da variável do tipo int e a do tipo hotdog criado por mim
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)
}
