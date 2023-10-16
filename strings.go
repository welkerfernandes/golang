package main

import (
	"fmt"
)

func main() {
	s := `Oi Welker e
		ta ok tudo bem`
	a := []byte(s)

	for _, v := range a {
		fmt.Printf("%v - %T - %#U - %#X\n", v, v, v, v)
	}

}
