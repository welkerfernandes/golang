package main

import (
	"fmt"
)

func main() {
	qalquer := map[int]string{
		123: "Muito Legal",
		498: "menos legal",
		985: "Esse é massa",
	}

	fmt.Println(qalquer)

	total := 0
	for i := range qalquer {
		total += i
	}

	fmt.Println(total)

	delete(qalquer, 498)

	fmt.Println(qalquer)
}
