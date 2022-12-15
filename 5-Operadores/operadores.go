package main

import "fmt"

func main() {
	// Aritméticos
	fmt.Println("***********************Aritméticos**********************")
	soma := 1 + 1
	subtracao := 5 - 1
	divisao := 10 / 2
	multiplicacao := 5 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 30
	var numero2 int32 = 20

	soma1 := numero1 + int16(numero2)
	fmt.Println(soma1)

	// Atribuição
	fmt.Println("\n***********************Atribuição***********************")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais
	fmt.Println("\n*****************Operadores Relacionais*****************")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	fmt.Println("\n*******************Operadores Lógicos*******************")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores Unários
	fmt.Println("\n*******************Operadores Unários*******************")
	numero3 := 10
	fmt.Println(numero3)
	numero3++
	fmt.Println(numero3)
	numero3 += 15
	fmt.Println(numero3)
	numero3--
	fmt.Println(numero3)
	numero3 -= 3
	fmt.Println(numero3)
	numero3 *= 2
	fmt.Println(numero3)
	numero3 /= 2
	fmt.Println(numero3)
	numero3 %= 3
	fmt.Println(numero3)

	// Operador Ternário (NÃO EXISTE EM GO)
	fmt.Println("\n********************Operador Ternário*******************")
	//txt := numero3 > 5 ? "Maior que 5" : "Menor que 5"

	var txt string
	if numero3 > 5 {
		txt = "Maior que 5"
	} else {
		txt = "Menor que 5"
	}
	fmt.Println(txt)

}
