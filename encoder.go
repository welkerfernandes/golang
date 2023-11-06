package main

import (
	"encoding/json"
	"os"
)

type carro struct {
	Marca    string
	Modelo   string
	Economia float64
	Ano      int
}

func main() {

	veiculo := carro{
		Marca:    "fiat",
		Modelo:   "Siena Fire",
		Economia: 1.0,
		Ano:      2010,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(veiculo)

}
