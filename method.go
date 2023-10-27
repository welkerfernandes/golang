package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) bomdia() {
	fmt.Println(p.nome, "Bom dia")
}

func main() {

	mauricio := pessoa{
		nome:  "mauricio",
		idade: 41}
	mauricio.bomdia()

}
