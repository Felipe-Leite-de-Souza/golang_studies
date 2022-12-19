package main

import "fmt"

// Por padrão o GO na estrutura do Switch não necessita de um Break pois o mesmo já existe internamente

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

// Outra maneira de se usar o Switch -> Quando é utilizada mais de uma variável
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

// Usando retorno com variável
func diaDaSemana3(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

// Fallthrough -> executa a próxima condição sem avaliar se a mesma é procedente
func diaDaSemana4(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

func main() {

	dia1 := diaDaSemana(3)
	fmt.Println(dia1)

	dia10 := diaDaSemana(8)
	fmt.Println(dia10)

	fmt.Println("-------------------------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
	dia20 := diaDaSemana2(22)
	fmt.Println(dia20)

	fmt.Println("-------------------------")
	dia3 := diaDaSemana3(5)
	fmt.Println(dia3)
	dia30 := diaDaSemana3(27)
	fmt.Println(dia30)

	fmt.Println("-------------------------")
	dia4 := diaDaSemana4(1)
	fmt.Println(dia4)
}
