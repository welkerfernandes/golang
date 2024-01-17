package main

//criando o proprio tipo
import (
	"fmt"
)

type hotdog int

var b hotdog = 209

func main() {
	x := 10

	fmt.Printf("B: %T X: %T\n\n\n", b, x)

	//conversao de tipo

	x = int(b)

	fmt.Printf("%v %T", x, x)
}
