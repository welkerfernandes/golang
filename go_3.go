package main

import (
	"fmt"
)

// variavel y está fora do codeblock por isso consigo repassar ela para todos os codeblock
var y = 10

func main() {

	qualquercoisa(y)
}

func qualquercoisa(x int) {

	fmt.Println(x)

}
