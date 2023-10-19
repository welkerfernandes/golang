package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "maça", "Jaca"}

	slice = append(slice, "melancia")
	for indice, valor := range slice {
		fmt.Println("No indice", indice, "Temos o valor", valor)
	}
}
