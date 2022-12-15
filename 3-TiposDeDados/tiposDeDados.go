package main

import (
	"errors"
	"fmt"
)

func main() {

	// Limite do int8 = 3 casas decimais
	// Aceita valores negativos
	var numeroInt8 int8 = 100
	fmt.Println("Número -> Int8:", numeroInt8)

	// Limite do int16 = 5 casas decimais
	// Aceita valores negativos
	var numeroInt16 int16 = 10000
	fmt.Println("Número -> Int16:", numeroInt16)

	// Limite do int32 = 10 casas decimais
	// Aceita valores negativos
	var numeroInt32 int32 = 1000000000
	fmt.Println("Número -> Int32:", numeroInt32)

	// Limite do int64 = 19 casas decimais
	// Aceita valores negativos
	var numeroInt64 int64 = 1000000000000000000
	fmt.Println("Número -> Int64:", numeroInt64)

	// Limite do int -> int64
	// Usa a arquitetura do computador como base
	// Aceita valores negativos
	var numeroInt int = 1234567890123456789
	fmt.Println("Número -> Int (arquitetura 64 bits = int64):", numeroInt)

	// uint -> unsigned int (int sem sinal)
	// O tamanho das casas decimais suportados são os mesmos dos int's normais acima
	var numeroUint8 uint8 = 123
	fmt.Println("Uint8: ", numeroUint8)

	var numeroUint16 uint16 = 12345
	fmt.Println("Uint16: ", numeroUint16)

	var numeroUint32 uint32 = 3332221110
	fmt.Println("Uint32: ", numeroUint32)

	var numeroUint64 uint64 = 9998887776665554443
	fmt.Println("Uint64: ", numeroUint64)

	var numeroUint uint = 9876543210987654321
	fmt.Println("Uint: ", numeroUint)

	// Alias
	// INT32 = rune
	var numeroRune rune = 987654321
	fmt.Println("Alias de int32 -> Rune:", numeroRune)

	// Alias
	// UINT8 = byte
	var numeroByte byte = 123
	fmt.Println("Alias de uint8 -> Byte:", numeroByte)

	var numeroFloat32 float32 = 250.55
	fmt.Println("Float32:", numeroFloat32)

	var numeroFloat64 float64 = 1230.73
	fmt.Println("Float64:", numeroFloat64)

	// Usa a arquitetura do computador como base
	numeroFloat := 123.45
	fmt.Println("Float -> arquitetura 64 bits:", numeroFloat)

	var str = "Texto"
	fmt.Println("str:", str)

	var str2 = "Texto2"
	fmt.Println("str2:", str2)

	// Número referente a tabela ASCII
	// Char é do tipo INT
	char := 'A'
	fmt.Println("Char 'A' -> tabela ASCII:", char)

	// Valor 0 (ZERO) - Ocorre quando a variável é inicializada e não é atribuida valor a ela.

	// Vazio
	var texto string
	fmt.Println("Texto (variável tipo string) inicializado sem atrubição:", texto)

	// Valor 0
	var tipoInt int16
	fmt.Println("Int (variável tipo int) sem atrubição:", tipoInt)

	// Valor false
	var booleano bool
	fmt.Println("Booleano (variável tipo boolean) sem atrubição:", booleano)

	// Valor declarado (true)
	var booleano1 bool = true
	fmt.Println("Booleano1:", booleano1)

	// Tipo error -> Valor zero = nil
	var erro error
	fmt.Println("Error (variável tipo error) sem declaração:", erro)

	// Erro declarado
	var erro1 error = errors.New("Erro interno")
	fmt.Println("Error declarado:", erro1)

}
