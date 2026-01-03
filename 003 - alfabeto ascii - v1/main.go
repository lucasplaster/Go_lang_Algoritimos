package main

import (
	"fmt"
	"os"
	"strings"


	"alfabeto_ascii/dados"
	"alfabeto_ascii/ferramentas"
)

func main() {

	if len(os.Args) > 1{
		var palavra string = strings.ToUpper(os.Args[1])

		palavra = ferramentas.RemoveAcentos(palavra)
		palavra = ferramentas.RemoveSimbolos(palavra)

		for linha := 0; linha < 5; linha++ {
			for _, ch := range palavra {
				letra := dados.Letras[ch]
				for coluna := 0; coluna < 5; coluna++ {
					if letra[linha][coluna] == 1 {
						fmt.Print("█")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Print("  ")
			}
			fmt.Println()
		}
	}else {
		fmt.Println("Use uma palavra como argumento.")
	}
	
	
}
