//problema de overflow tipo odometro de carro
//quando chega no limite 999 começa de novo 0 1 2 3

package main

import (
	"fmt"
)

var i uint16

func main() {
	i = 65535
	fmt.Println(i)

	i++
	fmt.Println(i)

}
