package main

import (
	"fmt"
)

type welker int

var x welker

func main() {

	fmt.Printf("%v %T\n\n", x, x)

	x = 42

	fmt.Printf("%v %T\n\n", x, x)

}
