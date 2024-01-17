package main

import (
	"fmt"
)

func main() {

	s := "Hello Playground"
	sb := []byte(s)

	for _, V := range sb {
		fmt.Printf("varlo: %v\n Tipo: %T\n %#U %#X\n\n\n\n", V, V, V, V)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v %T %#U %#X", s[i], s[i], s[i], s[i])
	}

} // final função main
