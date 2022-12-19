package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if init
	// variável limitada ao escopo do if
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Número é maior do que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor do que -10")
	} else {
		fmt.Println("Número entre 0 e -10")
	}
}
