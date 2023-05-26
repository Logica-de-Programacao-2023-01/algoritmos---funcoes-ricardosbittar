package main

import (
	"errors"
	"math"
)

func verificarPrimo(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("o número é negativo")
	}

	if numero < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}

	return true, nil
}
