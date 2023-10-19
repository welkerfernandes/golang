package main

import (
	"fmt"
)

// em go não existe while mas pode ser feito o for dessa forma descrita abaixo
func main() {

	x := 0

	for x <= 10 {
		fmt.Println("Você é vencedor")
		x++
	}

}
