package main

import (
	"fmt"
)

func main() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "welker",
		idade: 10,
	}

	fmt.Println(x)
}
