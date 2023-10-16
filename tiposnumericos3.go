package main

import (
	"fmt"
	"runtime"
)

func main() {

	//exibe o sistema operacional
	fmt.Println(runtime.GOOS)

	//exibe quantos bits tem o processador
	fmt.Println(runtime.GOARCH)
}
