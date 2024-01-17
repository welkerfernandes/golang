package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 20
	c := 30

	d := (a == b)
	e := (b <= a)
	f := (c >= b)

	fmt.Println("", a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)
}
