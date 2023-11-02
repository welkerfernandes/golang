package main

import (
	"fmt"
)

func main() {

	x := 0
	y := &x

	*y++

	fmt.Println(*y)
	fmt.Printf("%T,%T\n", x, y)
	fmt.Println(x, y)

}
