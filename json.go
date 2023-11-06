package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Nome           string
	Sobrenome      string
	Idade          int
	Profissao      string
	Conta_bancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:           "James",
		Sobrenome:      "Bond",
		Idade:          40,
		Profissao:      "Agente secreto",
		Conta_bancaria: 1000000.50,
	}

	darthvader := pessoa{
		Nome:           "Anakin",
		Sobrenome:      "Skywalker",
		Idade:          50,
		Profissao:      "Vilão Intergalactico",
		Conta_bancaria: 500000000000.83,
	}

	j, err := json.Marshal(jamesbond)

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}
