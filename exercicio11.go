package main

import (
	"fmt"
)

func main() {

	a := 150

	fmt.Println("binário\t\tdecimal\t\thexa-decimal")

	//binário decimal e hexadecimal
	fmt.Printf("%b\t\t", a)
	fmt.Printf("%d\t\t", a)
	fmt.Printf("%#x\t\t", a)

}
