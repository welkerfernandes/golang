package main

import (
	"fmt"
)

func calendario() {

	for ano := 0; ano <= 1; ano++ {
		fmt.Println("Ano ", ano)
		for meses := 1; meses <= 12; meses++ {
			fmt.Println("Meses ", meses)
		}
	}
}

func main() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora: ", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}
		fmt.Println("\n")
	}

	fmt.Println("=========================================")
	fmt.Println("Calendário")
	calendario()
}
