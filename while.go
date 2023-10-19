package main

import (
	"fmt"
)

// em go não existe while mas pode ser feito o for dessas forma descritas abaixo
func main() {

	x := 0

	for x <= 10 {
		fmt.Println("Você é vencedor")
		x++
	}

	//aqui utilizo for sem condição
	a := 0

	for {

		if a < 10 {
			a++
			fmt.Println(a)
		} else {
			break
		}

	}

}
