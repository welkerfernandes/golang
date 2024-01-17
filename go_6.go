package main

import (
	"fmt"
)

func main() {

	/*interpreted_literals := "Oi como vai você\n\n\n"
	fmt.Println(interpreted_literals)

	raw_literals := `"asds
	asdd
	sadasdas
	dsadsadsa
	"`
	fmt.Println(raw_literals)
	*/
	//diferença de Print Println Printf
	x := "X Oi bom dia"
	y := "Y Bom dia"
	z := fmt.Sprint(x, " ", y)
	fmt.Println(z)

}
