package main

import (
	"fmt"
)

func main() {
	// // var colours map[string]string
	// colours := make(map[string]string)
	// colours["white"] = "#ffffff"
	// delete(colours, "white")

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex value for", colour, "is", hex)
	}
}
