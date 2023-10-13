package main

import (
	"fmt"
)

// package level scope que fica atribuido acima do escopo da função atribuida nas funções
// para usar o var precisa ser fora do codeblock
// não posso usar gopher := fora do codeblock por isso utiçizei =
var y = 13

func main() {
	//aqui declarei uma variável e depois atribui ela em uma função chamada qualquer coisa
	y := 4
	qualquercoisa(y)
}

// aqui pego a função qualquer coisa e executo a ação que eu quero com base no valor atrubuido na função main acima que foi aonde a variável foi declarada
func qualquercoisa(x int) {
	fmt.Println(y, ": package level scope") // essa variável não foi declada dentro do codeblock mas irei declarar com var lá em cima
	fmt.Println(x, ": var atribuida na func main")
}
