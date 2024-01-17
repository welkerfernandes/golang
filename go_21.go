package main

import "fmt"

const (
	_ = iota
	P = 15
	S
	T
)

func main() {
	fmt.Printf("decimal: %d \n binário: %b \n Hexadecimal: %#X\n\n", P, P, P, S, S, S, T, T, T)
}
