package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {
		//continue é usado para quebrar execução da iteração
		if i == 6 {
			continue
		}
		fmt.Println(i)
		fmt.Println("Sai da iteração e vou iniciar a próxima iteração dentro do laço")
	}

	//ja o break quebra o loop em si
	fmt.Println("Inicia outo laço for para mostrar o break")
	a := 0
	for {
		if a <= 10 {
			a++
			fmt.Println(a)
		} else {
			break
		}
	}

	fmt.Print("Sai do laço FOR")
}
