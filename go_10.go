package main

import (
	"fmt"
)

var x int = 42
var y string = "James bond"
var z bool = true

func main() {

	s := fmt.Sprintf("|%v %T |%v %T |%v %T", x, x, y, y, z, z)
	fmt.Println(s)

}
