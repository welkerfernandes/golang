package main

import (
	"fmt"
)

func main() {

	for mes := 1; mes <= 12; mes++ {

		fmt.Println("mes", mes)
		for y := 1; y < 31; y++ {

			fmt.Println("dia", y)

		}
		fmt.Println("\n")
	} //horas

} //func main
