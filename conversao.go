package main

import (
	"fmt"
)

// aqui tem o tipo criado por mim chamado hot dog
type hotdog int

// aqui atibui o valor ao tipo hotdog e a variável b
var b hotdog = 10

func main() {
	//aqui atribuí uma variável tipo int
	x := 10

	//no fim ambas são int e mostram o mesmo valor, vamos ealizar uma conversão de tipos agora.
	fmt.Printf("%v\n", x)

	//aqui convertemos uma para o tipo da outra
	x = int(b)
	fmt.Printf("%v\n", x)

}
