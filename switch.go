package main

import (
	"fmt"
)

func main() {

	x := 10

	switch true {

	case x > 10:
		fmt.Println("maior que 10")
	case x == 50:
		fmt.Println("igual a 50")
	case x < 20:
		fmt.Println("menor que 20")

	}

} //fim main
