package gerador

import "ordinais/dados"

func Ordinal(n int)string{
	if n == 0{
		return ""
	}

	centena := n / 100
	dezena := n % 100 / 10
	unidade := n % 10
	
	if n < 10 {
		return dados.Unidades[unidade]
	}

	if n < 100 {
		if unidade == 0 {
			return dados.Dezenas[dezena]
		}
		return dados.Dezenas[dezena] + " " + dados.Unidades[unidade]
	}

	resto := n%100

	if resto == 0 {
		return dados.Centenas[centena]
	}
	return dados.Centenas[centena] + " " + Ordinal(resto)
}