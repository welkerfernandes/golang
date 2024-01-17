package main

import (
	"fmt"
)

type welker int

var x welker
var b int

func main() {
	fmt.Printf("%v %T", x, x)
	x = 42
	fmt.Printf("%v %T", x, x)

	b = int(x)
	fmt.Printf("\n\n\n%v %T", b, b)
}
