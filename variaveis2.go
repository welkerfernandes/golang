package main

import (
	"fmt"
)

func main() {

	//em go existem keywords que já vem pré configuradas e não podem ser usadas como nome de variável como package eu não posso criar uma variável com esse nome por exemplo
	//abaixo vemos que o go já faz o calculo com aqueles valores na variável x
	x := 10 + 10
	fmt.Println("\n", x, "\n")

	// e também mostra valores bool
	//statemant pega o valor já salva o valor do resultado sendo uma declaração
	y := 10 > 20
	fmt.Println(y)

}
