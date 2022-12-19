package main

import (
	"fmt"
)

// Só pode existir um parâmetro variático na função
// O parâmetro variático é sempre o último parâmetro da função
func soma(numeros ...int) int {
	total := 0
	for _, nunumero := range numeros {
		total += nunumero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, nunumero := range numeros {
		fmt.Println(texto, nunumero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4)
	fmt.Println(totalDaSoma)
	escrever("Hello Word!", 1, 2, 3, 4)
}
