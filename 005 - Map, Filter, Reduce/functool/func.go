package functool

func Filter[T any](valores []T, fn func(T) bool) []T {
	resultado := make([]T, 0, len(valores))

	for _, valor := range valores {
		if fn(valor) {
			resultado = append(resultado, valor)
		}
	}

	return resultado
}

func Map[T any, R any](valores []T, fn func(T) R) []R {
	resultado := make([]R, len(valores))

	for indice, valor := range valores {
		resultado[indice] = fn(valor)
	}

	return resultado
}

func Reduce[T any](valores []T, fn func(T, T) T, inicio T) T {
	var resultado T = inicio

	for _, valor := range valores {
		resultado = fn(resultado, valor)
	}
	return resultado
}
