package main

import "fmt"

// NÃO EXISTE HERANÇA EM GO

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Helena", "Souza", 23, 160}
	fmt.Println(p1)

	e1 := estudante{p1, "Cyber Security", "USCS"}
	fmt.Println(e1.nome)
}
