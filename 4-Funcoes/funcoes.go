package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadosSoma, resultadosSubtracao := calculosMatematicos(20, 5)
	fmt.Println(resultadosSoma, resultadosSubtracao)

	resultadosSoma1, _ := calculosMatematicos(30, 20)
	fmt.Println(resultadosSoma1)
}
