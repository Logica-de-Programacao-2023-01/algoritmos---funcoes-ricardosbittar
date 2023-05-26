package main

import (
	"errors"
	"sort"
)

func ordenarValores(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	sliceOrdenado := make([]int, len(slice))
	copy(sliceOrdenado, slice)

	sort.Ints(sliceOrdenado)

	return sliceOrdenado, nil
}
