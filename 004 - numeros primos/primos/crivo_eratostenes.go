package primos

func NumerosPrimosCrivo(numeroMaximo int) []int {
	// https://pt.wikipedia.org/wiki/Crivo_de_Erat%C3%B3stenes
	if numeroMaximo < 2 {
		return []int{}
	}
	primos := make([]bool, numeroMaximo+1)

	for i := range primos {
		primos[i] = true
	}

	primos[0] = false
	primos[1] = false

	cont := 2
	for cont*cont <= numeroMaximo {
		if primos[cont] {
			for j := cont * cont; j <= numeroMaximo; j += cont {
				primos[j] = false
			}
		}
		cont++
	}

	resultado := []int{}
	for i := 2; i <= numeroMaximo; i++ {
		if primos[i] {
			resultado = append(resultado, i)
		}
	}

	return resultado
}
