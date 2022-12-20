package main

import "fmt"

// Panic -> Mata a execução do programa
// Recover -> Recupera do a execução do programa após o Panic

func recuoerarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução!")
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuoerarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")

}
