package primos

import "math"

func NumerosPrimosRaiz(numeroMaximo int) []int {
	if numeroMaximo < 2 {
		return []int{}
	}

	var primos = []int{}
	var primo bool

	for numero := 2; numero < numeroMaximo; numero++ {
		primo = true

		for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
			if numero%i == 0 {
				primo = false
				break
			}
		}

		if primo {
			primos = append(primos, numero)
		}
	}
	return primos
}
