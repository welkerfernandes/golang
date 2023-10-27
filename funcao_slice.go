package main

import (
	"fmt"
)

func main() {

	si := []int{10, 10, 1, 2, 3, 4, 5}

	total := calculo(si...)
	fmt.Println(total)
}

func calculo(x ...int) int {
	calculo := 0
	for _, v := range x {
		calculo += v
	}
	return calculo
}
