package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Amanda",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Letícia",
			"ultimo":   "Perreira",
		},
		"curso": {
			"nome":   "Química",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["sexo"] = map[string]string{
		"sexo": "Feminino",
	}
	fmt.Println(usuario2)
}
