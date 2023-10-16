package main

import (
	"fmt"
	"runtime"
)

func main() {

	x := 10

	y := float64(x)
	y = 64.45

	fmt.Printf("%v = %T\n", y, y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
