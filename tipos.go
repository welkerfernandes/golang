package main

import (
	"fmt"
)

// se eu declarar uma variavel como int ela será int até o fim do programa
// se precisar de outro tipo então precisará de outra variável
// go é linguagem de tipagem estática
//se eu declarar o tipo o go vai identificar sozinho

var x int = 10

// também posso fazer o tipo de uma variável e fazer sua declaração dentro do codeblock sem usar gopher
var z int

func main() {

	//nesse caso eu dei um novo valor a variável x sem alterar seu tipo declarado na variavel x do package level scope
	//se eu tentar declarar x como float 20.3 daria erro pois no package level scope x eu declarei como int
	x := 20
	//fmt.Println(x)

	// aqui com o print format (printf) posso ver o formato declarado do dado

	fmt.Printf("%v , %T\n", x, x)

	z = 50
	fmt.Printf("\n\n%v, %T", z, z)
}
