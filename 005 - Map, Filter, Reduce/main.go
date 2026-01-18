package main

import (
	"fmt"
	"highOrderFunction/functool"
)

var valores = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	pares := functool.Filter(
		valores,
		func(valor int) bool {
			return valor%2 == 0
		})

	dobro := functool.Map(
		valores,
		func(valor int) int {
			return valor * 2
		},
	)

	soma := functool.Reduce(
		valores,
		func(acumulador, valor int) int {
			return acumulador + valor
		},
		0,
	)

	fmt.Println(valores)
	fmt.Println(pares)
	fmt.Println(dobro)
	fmt.Println(soma)
}
