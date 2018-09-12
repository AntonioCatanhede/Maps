package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":    "#FF0000",
		"blue":   "#0000FF",
		"green":  "#008000",
		"yellow": "#FFFF00",
	}
	// modificando valores
	colors["red"] = "Minha cor preferida"
	//adicionando valores
	colors["pink"] = "#FFC0CB"
	//deletando pares chave-valor
	delete(colors, "pink")
	printMap(colors)

}

func printMap(m map[string]string) {

	for key, value := range m {
		fmt.Println(key, "é representado por", value)
	}
	//maneira alternativa
	for key := range m {
		fmt.Println(key, "é representado por", m[key])
	}

}
