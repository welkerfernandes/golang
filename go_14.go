package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := "e"
	b := "é"
	fmt.Printf("%v %T \n %v %T\n \n \n ", a, a, b, b)

	d := []byte(a)
	e := []byte(b)

	fmt.Printf("%v %T \n %v %T\n\n\n\n", d, d, e, e)

	y := 10
	x := 10.0

	fmt.Printf("%v %T \n\n %v %T\n\n\n\n\n\n", y, y, x, x)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
