package main

import "errors"

func verificarPresenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("o slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}
