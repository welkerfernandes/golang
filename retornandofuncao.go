package main

import (
	"fmt"
)

func main() {
	x := retornafuncao()
	y := x(3)
	fmt.Println(y)
}

func retornafuncao() func(int) int {

	return func(i int) int {
		return i * 10
	}
}
