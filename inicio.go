package main

import (
	"fmt"
)

func main() {
	//aqui começa o programa
	//pacote . identificador
	// estou usando o _, no inicio para eliminar uma variável vazia pois no golang é proibida variavel vazia
	_, erros := fmt.Println("Hello World", "Oi Galera", 100)
	fmt.Println(erros)
	//aqui acaba seu programa

}
