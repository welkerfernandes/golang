// loop while não existe no golang
// while funciona da seguinte forma caso algo for true realiza uma ação
package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 10 {
		fmt.Println("Está correto")
		x++
	}

	y := 0

	for {
		if y <= 10 {
			y++
			fmt.Println("\n\n\ny maior que 10 to fora")
		} else {
			fmt.Println("\n\n\ny ta errado")
			break
		}
	} //fim do for
}
