package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)  // referência
	fmt.Println(variavel3, *ponteiro) // desferênciação

	variavel3 = 101
	fmt.Println(variavel3, ponteiro)  // referência
	fmt.Println(variavel3, *ponteiro) // desferênciação

}
