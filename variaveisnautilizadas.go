package main

import (
	"fmt"
)

func main() {
	_, erros := fmt.Println("Hello world", "nome", 100)
	fmt.Println(erros)
}
