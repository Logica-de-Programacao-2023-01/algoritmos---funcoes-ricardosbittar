package main

import (
	"errors"
	"strings"
	"unicode"
)

func filtrarStringsComMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	var resultBuilder strings.Builder

	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			resultBuilder.WriteString(str)
			resultBuilder.WriteString(" ")
		}
	}

	result := resultBuilder.String()
	if len(result) > 0 {
		result = strings.TrimRight(result, " ")
	}

	return result, nil
}
