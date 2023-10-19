package main

import (
	"fmt"
)

var x [5]int

func main() {
	x[0] = 1
	x[3] = 3
	fmt.Println(x)
}
