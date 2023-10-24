package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		//indice 0  1  2  3  4  5
		{1, 2, 3, 4, 5, 6},       //0
		{7, 8, 9, 10, 11, 12},    //1
		{13, 14, 15, 16, 17, 18}, //2
	}

	fmt.Println(ss[1][1])
}
