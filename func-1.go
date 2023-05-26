package main

func media(valores []int) float64 {
	soma := 0
	for _, v := range valores {
		soma += v
	}

	media := float64(soma) / float64(len(valores))

	return media
}
