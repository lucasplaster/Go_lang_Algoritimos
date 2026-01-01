package main

import (
	"fmt"
	"os"
	"strconv"

	"ordinais/gerador"
)




func main(){
	if len(os.Args) < 2{
		fmt.Println("Você deve usar ordinais <numero> ou ordinais <num1> <num2> <num3> ...")
		return
	}

	numerosOrigem := os.Args[1:]

	for indice, valor := range numerosOrigem{
		valor64, err := strconv.ParseInt(valor, 10, 64)

		if err != nil{
			fmt.Printf("%s não é um numero válido na posição %d.", valor, indice)
		}
		
		valor := int(valor64)
		fmt.Printf("%d = %s\n", valor, gerador.Ordinal(valor))


	}
}