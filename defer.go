package main

import (
	"fmt"
)

func main() {
	//defer é tipo deixar para última hora
	defer fmt.Println("COM Defer")
	fmt.Println("Sem Defer")

}
