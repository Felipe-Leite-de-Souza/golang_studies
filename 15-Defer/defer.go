package main

import "fmt"

// DEFER = ADIAR
// Muito utilizado com a utilização de banco de dados para fechar a conexão

func func1() {
	fmt.Println("Func 1")
}

func func2() {
	fmt.Println("Func 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	func1()
	func2()

	defer fmt.Println(alunoEstaAprovado(7, 8))
}
