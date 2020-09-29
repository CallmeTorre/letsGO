package main

import "fmt"

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println(color, hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
	fmt.Println(colors)
}
