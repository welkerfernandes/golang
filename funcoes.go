package main

import (
	"fmt"
)

func main() {
	valor := soma(10, 10)
	fmt.Println(valor, "\n\n")
	basica()
	argumentos("tarde")

	total, quantos := calculo(10, 20, 1, 2, 3)
	fmt.Println(total, quantos)
}

func basica() {
	fmt.Println("oi Bom dia")
}

func argumentos(s string) {
	if s == "manhã" {
		fmt.Println("Bom dia")
	} else if s == "tarde" {
		fmt.Println("Boa tarde")
	} else {
		fmt.Println("Olá")
	}
}

func soma(x, y int) int {

	return x + y
}

func calculo(x ...int) (int, int) {
	calculo := 0
	for _, v := range x {
		calculo += v
	}
	return calculo, len(x)
}
