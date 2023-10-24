package main

import (
	"fmt"
	"sync"
)

func nomes(wg *sync.WaitGroup) {

	defer wg.Done()
	paises := []string{"Brasil", "França", "Israel", "Estados Unidos"}

	for _, valor := range paises {
		fmt.Println(valor)
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go nomes(&wg)

	fmt.Println("Pronto")

	wg.Wait()
}
