/*
o break quebra o for inteiro
e o continue para a iteração do loop
*/
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {

		if i%2 != 0 {
			//faz isso quando o numero é impar
			continue

		}
		fmt.Println(i)

	}
}
