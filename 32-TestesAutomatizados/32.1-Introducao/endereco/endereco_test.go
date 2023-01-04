// TESTE UNITÁRIO BÁSICO

package endereco

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada ABC", "Estrada"},
		{"Praça das Rpsas", "Tipo inválido"},
		{"RUA Teste", "Rua"},
		{"AVENIDA Teste", "Avenida"},
		{"", "Tipo inválido"},
	}

	//enderecoParaTeste := "Rua ABC"
	//tipoDeEnderecoEsperado := "Rua"
	//tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//	t.Errorf("O tipo é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	//}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
