// TESTE UNITÁRIO BÁSICO

package endereco

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua ABC"

	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
