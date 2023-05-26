package main

import (
	"errors"
	"strconv"
)

func somarDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("o número é negativo")
	}

	numeroStr := strconv.Itoa(numero)
	soma := 0

	for _, char := range numeroStr {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, err
		}
		soma += digito
	}

	return soma, nil
}
