package main

import "fmt"

// A função init pode existir uma por arquivo
// Executa sempre antes das demais execuções
// Utilizada para iniciar alguns valores ou contruir alguns setup's

var n int

func init() {
	fmt.Println("Executando a função init!")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executanda!")
	fmt.Println(n)
}
