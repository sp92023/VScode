package main

import "fmt"

func main() {
	/*colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}*/

	//var colors map[string]string = colors := make(map[string]string)
	/*colors := make(map[string]string)
	colors["white"] = "#ffffff"
	fmt.Println(colors)*/

	/*colors := make(map[int]string)
	colors[10] = "#ffffff"

	delete(colors, 10)
	fmt.Println(colors)*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
