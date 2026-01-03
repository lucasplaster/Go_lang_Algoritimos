package ferramentas

import (
	"strings"

	"alfabeto_ascii/dados"
)


func RemoveAcentos (s string) string{
	stringFormatada := strings.Builder{}

	for _, r := range s {
		if valor, existe := dados.MapaAcentuacao[r]; existe {
			stringFormatada.WriteRune(valor)
		}else{
			stringFormatada.WriteRune(r)
		}
	}
	return stringFormatada.String()
}

func RemoveSimbolos(s string) string{
	stringFormatada := strings.Builder{}

	for _, r := range s {
		if (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') {
			stringFormatada.WriteRune(r)
		}
	}

	return stringFormatada.String()
}