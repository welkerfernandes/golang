package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := cliente{
		nome:      "Welker",
		sobrenome: "Fernandes",
		fumante:   false,
	}

	cliente2 := cliente{
		"Daniela",
		"Abrenhosa",
		false}

	fmt.Println(cliente1, cliente2)
}
