package main

import (
	"fmt"
	"primos/primos"
)

func main() {
	fmt.Println("Usando metodo do divisor")
	fmt.Println(primos.NumerosPrimosSimples(100))
	fmt.Println("Usando metodo da raiz")
	fmt.Println(primos.NumerosPrimosRaiz(100))
	fmt.Println("Usando metodo do crivo de eratostenes")
	fmt.Println(primos.NumerosPrimosCrivo(100))
	fmt.Println()
	fmt.Println("Verificando se 100 é primo: ", primos.IsPrimo(100))
	fmt.Println("Verificando se 101 é primo: ", primos.IsPrimo(101))

}
