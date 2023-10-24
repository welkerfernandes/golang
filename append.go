package main

import (
	"fmt"
)

func main() {

	slices := []int{1, 5, 7, 8, 6, 5}
	outra_slice := []int{1, 2, 3, 4, 5, 6, 7}

	slices = append(slices, 5, 6, 7, 8)

	fmt.Println(slices)

	slices = append(slices, outra_slice...)

	fmt.Println(slices)

}
