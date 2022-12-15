package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Struct")

	var u usuario
	u.nome = "Fernada"
	u.idade = 30
	fmt.Println(u)

	enderecoTest := endereco{"Rua A", 10}
	u2 := usuario{"Bianca", 20, enderecoTest}
	fmt.Println(u2)

	u3 := usuario{nome: "Rebeca"}
	fmt.Println(u3)
}
