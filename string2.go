package main

import (
	"fmt"
)

func main() {

	s := "Welker é o Brabo"

	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#X\n", v, v, v, v)
	}

	fmt.Println("")
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}
}
