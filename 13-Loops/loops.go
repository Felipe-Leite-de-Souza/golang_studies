package main

import (
	"fmt"
	"time"
)

func main() {

	// Em GO não é possível fazer range em Structs
	i := 0

	// Variável declarada fora do escopo do for
	// Imprime até 10
	for i < 10 {
		i++
		fmt.Println("Valor de i:", i)
		time.Sleep(time.Second)
	}

	fmt.Println("---------------")

	// Variável declarada dentro do escopo do for
	// Imprime até 9
	for j := 0; j < 10; j++ {
		fmt.Println("Valor de j:", j)
		time.Sleep(time.Second)
	}

	// Range
	nomes := [3]string{"Abacaxi", "Banana", "Caju"}

	// i -> índice
	// n -> nome
	// _ -> ignorar parâmetro
	for _, n := range nomes {
		fmt.Printf("Fruta: %s\n", n)
	}

	fmt.Println("---------------")

	// Iterando string
	for i, l := range "TesTES" {
		fmt.Printf("Índice: %d -> Letra (Tabela ASCII): %v\n", i, l)
		fmt.Printf("Índice: %d -> Letra (Tabela ASCII): %v\n", i, string(l))
		fmt.Println()
	}

	fmt.Println()
	// Iterando map
	usuario := map[string]string{
		"nome":      "Brian",
		"sobrenome": "Huges",
	}

	fmt.Println()
	fmt.Println(usuario)

	for c, v := range usuario {
		fmt.Println(c, v)
	}

	// Loop infinito
	//for {
	//	fmt.Println("Loop infinito")
	//	time.Sleep(time.Second)
	//}
}
