package main

func segundoMaiorValor(numeros []int) int {

	maior := numeros[0]
	segMaior := numeros[1]

	if segMaior > maior {
		maior, segMaior = segMaior, maior
	}

	for i := 2; i < len(numeros); i++ {
		if numeros[i] > maior {
			segMaior = maior
			maior = numeros[i]
		} else if numeros[i] > segMaior {
			segMaior = numeros[i]
		}
	}

	return segMaior
}
