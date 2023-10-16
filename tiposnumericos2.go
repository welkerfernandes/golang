package main

import (
	"fmt"
)

// declarei um valor int
var x = 10

func main() {

	//fiz a conversão da variavel x de int para float64
	a := float64(x)
	//atribuí um valor float nela
	a = 20.85
	//exibi o valor float com seu tipo para confirmação da conversão
	fmt.Printf("%v = %T", a, a)

}
