package primos


func NumerosPrimosSimples(numeroMaximo int) []int {
	if numeroMaximo < 2 {
		return []int{}
	}

	divisor := 0
	primos := []int{}

	for i:=1 ; i <= numeroMaximo; i++ {
		divisor = 0

		for j :=1; j < i+1; j++ {
			if i % j == 0 {
				divisor += 1
			}
		}

		if divisor == 2 {
			primos = append(primos, i)
		}

	}

	return primos

}
