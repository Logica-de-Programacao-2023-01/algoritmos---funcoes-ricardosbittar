package main

import (
	"errors"
	"sort"
	"strings"
)

func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	// Ordena o slice em ordem alfabética
	sort.Strings(slice)

	// Junta as strings em uma única string separada por vírgulas
	result := strings.Join(slice, ",")

	return result, nil
}
