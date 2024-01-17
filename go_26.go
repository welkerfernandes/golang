package main

import (
	"fmt"
)

const (
	_    = iota
	Pano = (1994 + iota)
	Sano
	Tano
	Qano
)

func main() {

	fmt.Println(" Ano: ", Pano, "\n Ano: ", Sano, "\n Ano: ", Tano, "\n Ano: ", Qano)

}
