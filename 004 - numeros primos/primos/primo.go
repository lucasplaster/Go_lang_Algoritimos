package primos

import "math"

func IsPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}
