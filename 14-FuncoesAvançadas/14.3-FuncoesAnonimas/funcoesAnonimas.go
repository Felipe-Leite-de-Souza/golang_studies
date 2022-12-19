package main

import "fmt"

func main() {

	// Sem parâmetros
	func() {
		fmt.Println("Hello World!")
	}()

	// Com parâmetros
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	// Com retorno -> Não retorna nada pois não tem onde armazenar o conteúdo
	func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	// Com retorno -> Armazenado em uma variável
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println((retorno))
}
