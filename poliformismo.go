package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentes_arrancados int
	salario           float64
}

type arquiteto struct {
	pessoa
	tipo_de_construcao string
	tamanho_da_loucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "eu ja arranquei", x.dentes_arrancados, "e meu salário é de ", x.salario)
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "Faço constuções do tipo", x.tipo_de_construcao, " e minha maior obra é ", x.tamanho_da_loucura)
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Josefina",
			sobrenome: "Do Carmo",
			idade:     29,
		},
		tipo_de_construcao: "Prédios",
		tamanho_da_loucura: "Prédio transversal",
	}

	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Afonso",
			sobrenome: "Padilha",
			idade:     31,
		},
		dentes_arrancados: 7864,
		salario:           5200,
	}

	//mrdente.oibomdia()
	//mrpredio.oibomdia()

	serhumano(mrdente)
	serhumano(mrpredio)

} //fim do main
