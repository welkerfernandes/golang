package main

import (
	"fmt"
)

func main() {

	t := somente_pares(soma, []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somente_pares(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
