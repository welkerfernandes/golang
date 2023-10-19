package main

import (
	"fmt"
)

func fatia() {
	nomes := []string{"Welker", "Daniela", "Onofre", "Alice"}
	fatia := nomes[:]
	fmt.Println(fatia)
}

func main() {

	carros := []string{"Honda Civic", "Corolla", "HB20", "Siena"}

	for i := 0; i < len(carros); i++ {
		fmt.Println(carros[i])
	}

	fatia()

}
