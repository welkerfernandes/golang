package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"Alfredo": 5551234,
		"Welker":  555456,
	}

	amigos["gopher"] = 44444

	sera, ok := amigos["fantasma"]

	fmt.Println(sera, ok)

}
