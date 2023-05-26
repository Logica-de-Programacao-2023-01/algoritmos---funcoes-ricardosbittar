package main

import (
	"errors"
	"strings"
)

func substituirVogaisPorAsterisco(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("a string está vazia")
	}

	vogais := "aeiouAEIOU"
	resultado := strings.Builder{}

	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			resultado.WriteRune('*')
		} else {
			resultado.WriteRune(char)
		}
	}

	return resultado.String(), nil
}
