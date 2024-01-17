// deslocamento de bits
// deslocamos bits binarios da direita para esquerda
// bitwise operation
package main

import (
	"fmt"
)

const (
	_  = iota // 0
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	y := 24
	x := y << 2
	z := y >> 2

	fmt.Printf("x: %b\n y: %b\n z: %b\n\n\n\n", x, y, z)
	fmt.Printf("%b \n%d\n%b\n%d\n%b\n%d\n%b\n%d\n\n", KB, KB, MB, MB, GB, GB, TB, TB)
}
