package main

import (
	"fmt"
	"os"
	"strconv"
	"quicksort/quicksort"
)


func main(){
		if len(os.Args) <=1{
		fmt.Println("Numero de argumentos insuficientes.")
		os.Exit(-1)
	}

	entrada := os.Args[1:]

	numeros := make([]int, len(entrada))

	for indice, valor := range entrada {
		numero, erro := strconv.Atoi(valor)
		if erro != nil {
			fmt.Printf("o item [%s] na [%dª] posição não é um numero válido. ", valor, indice+1)
			os.Exit(-1)
		}
		numeros[indice] = numero
	}

	fmt.Println(quicksort.Quicksort(numeros))
}