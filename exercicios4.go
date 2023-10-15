package main

import (
	"fmt"
)

type welker int

var x welker

func main() {

	fmt.Printf("%v\n, %T\n", x, x)
	x := 42
	fmt.Println(x)
}
