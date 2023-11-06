package main

import (
	"encoding/json"
	"fmt"
)

type informacoes struct {
	Nome           string  `json:"Nome"`
	Sobrenome      string  `json:"Sobrenome"`
	Idade          int     `json:"Idade"`
	Profissao      string  `json:"Profissao"`
	Conta_bancaria float64 `json:"Conta_bancaria"`
}

func main() {

	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente secreto","Conta_bancaria":1000000.5}
`)

	var info informacoes

	err := json.Unmarshal(sb, &info)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(info)

}
