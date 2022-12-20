package main

import "fmt"

// Deve-se ter cuidado ao utilizar o tipo genérico pois abrirá mão dos tipos
// Qualquer tipo de parâmetro satisfaz a essa interface
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
	generica(33.4)

	// Um exemplo de interface genérica é o Println
	// Recebe várias interfaces do tipo que desejar
	fmt.Println(1, 2, "String", false, true, float64(12345))

	// Gera confusão pois não tem tipos definidos
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}
	fmt.Println(mapa)
}
