package main

import "fmt"

// Método -> Codependente de alguma estrutura

type usuario struct {
	nome  string
	idade uint8
}

// Passando o valor por cópia -> Qualquer alteração existe aoenas dentro do método
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados.\n", u.nome)
}

// Passando o valor por cópia -> Qualquer alteração existe aoenas dentro do método
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Passando o valor por ponteiro, para alterar o valor da idade do usuário
// Persiste a alteração
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Caroline", 30}
	usuario2.salvar()

	maiordeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiordeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
