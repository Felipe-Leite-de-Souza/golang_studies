package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string
	fmt.Println(array1)

	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 21)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "F"
	fmt.Println(slice2)

	// Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 6)
	slice3 = append(slice3, 7)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 13)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
