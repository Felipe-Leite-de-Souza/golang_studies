package main

import "fmt"

// Closure -> Funções que referenciam variáveis que estão fora do seu corpo de escopo

func closure() func() {
	texto := "Dentro da função closure!"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main!"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
