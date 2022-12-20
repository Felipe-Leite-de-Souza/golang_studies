package main

import "fmt"

// Não é recomendado utilizar quando se tem muitas execuções

// Informar quando deve ser parada para não gerar um Stack Overflow (Estouro de pilha)
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(12)
	fmt.Println(fibonacci(posicao))

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
