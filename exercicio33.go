package main

import (
	"fmt"
)

const x int = 10
const y = 10

func main() {

	fmt.Printf("%v - %v: %T - %T", x, y, x, y)
}
