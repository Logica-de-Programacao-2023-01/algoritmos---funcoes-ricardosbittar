package main

import (
	"errors"
)

func aplicarFuncao(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("o slice está vazio")
	}

	resultado := 0

	for _, valor := range slice {
		resultado += funcao(valor)
	}

	return resultado, nil
}
