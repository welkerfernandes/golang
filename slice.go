package main

import (
	"fmt"
)

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	//com slice consigo adicionar e com array não.
	slice2 := append(slice, 7, 8, 9)
	fmt.Println(slice2)
}
