package quicksort

func Particionar(numeros []int, pivo int) (menores []int, maiores []int){
	// Foi criado essa função para deixar a função principal mais limpa
	for _, valor := range numeros {
		if valor <= pivo {
			menores = append(menores, valor)
		}else{
			maiores = append(maiores, valor)
		}
	}

	return menores, maiores
}


func Quicksort(numeros []int) []int{
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indice := len(n) / 2
	pivo := n[indice]

	n = append(n[:indice], n[indice+1:]...)

	menores, maiores := Particionar(n, pivo)

	return append(
		append(Quicksort(menores), pivo), Quicksort(maiores)...
	)
}
