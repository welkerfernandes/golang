package main

import (
	"fmt"
)

func main() {
	//variável com um tipo
	var x uint16

	//coloquei o valor limite que é permitido pelo tipo
	x = 65535

	//aqui vou exibir o valor
	fmt.Println(x)

	//aqui vou adicionar um valor a mais para passar do limite permitido que é d e65535
	x++
	fmt.Println(x)

	//aqui vou adicionar um valor a mais para passar do limite permitido que é d e65535
	x++
	fmt.Println(x)
	//isso da um problema enorme em que é tipo um odometro do carro que quando chega a 9999 ele zera e começa tudo de novo
	//ocorreu um problema de um satélite da nasa que ocorreu algo parecido e acabou explodindo por conta desse problema
}
