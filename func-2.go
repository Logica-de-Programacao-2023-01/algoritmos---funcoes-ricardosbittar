package main

func vogais(valores []string) []string {
	vogaisOnly := []string{}

	for _, v := range valores {
		if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" {
			vogaisOnly = append(vogaisOnly, v)
		}

	}
	return vogaisOnly

}
