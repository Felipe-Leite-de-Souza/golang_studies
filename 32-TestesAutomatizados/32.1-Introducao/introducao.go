package main

import (
	"fmt"
	"introducaoTestes/endereco"
)

func main() {
	tipoDeEndereco := endereco.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoDeEndereco)
}
