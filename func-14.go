package main

import "errors"

func intersecaoSlices(slice1 []int, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("um dos slices está vazio")
	}

	intersecao := []int{}

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				intersecao = append(intersecao, num1)
				break
			}
		}
	}

	return intersecao, nil
}
