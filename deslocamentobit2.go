package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
)

func main() {

	fmt.Println("Binário \t\t\t\t\t\t\t Decimal")
	fmt.Printf("%b\t\t\t\t\t\t\t", KB)
	fmt.Printf("\t%d", KB)

	fmt.Printf("\n%b\t\t\t\t\t", MB)
	fmt.Printf("%d", MB)

	fmt.Printf("\n%b\t\t\t", GB)
	fmt.Printf("%d", GB)
}
