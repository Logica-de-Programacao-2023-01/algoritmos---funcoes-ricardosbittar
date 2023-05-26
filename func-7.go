package main

import "errors"

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	resultados := make([]int, len(slice))
	for i, valor := range slice {
		resultados[i] = funcao(valor)
	}

	return resultados, nil
}
