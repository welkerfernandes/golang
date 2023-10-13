package main

import (
	"fmt"
)

func main() {
	//dois tipos de prints
	x := "Bom dia\ncomo vai?\n"       //interpreter literal
	y := `Aqui tem a string cru \n\n` //raw strings
	z := fmt.Sprint(x, y)
	fmt.Println(z)
	//Print= mostra o resultado sem uma linha nova em branco abaixo
	//Println= mostra o resultado com uma linha em branco abaixo
	//Sprint=retorna o resultado como string sem colocar mais nenhuma informação na tela

}
