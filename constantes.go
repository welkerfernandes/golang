package main

import (
	"fmt"
)

// nesse caso a constate é permanente é bom utilizar const para nomes esáticos que não irão mudar
const (
	a = 10
	b = 50
	c = 45.5
)

func main() {

	fmt.Println(a, b, c)
}
