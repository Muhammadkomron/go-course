package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"violet": "#9400D3",
		"indigo": "#4B0082",
		"blue":   "#0000FF",
		"green":  "#008000",
		"yellow": "#FFFF00",
		"orange": "#FF7F00",
		"red":    "#FF0000",
	}
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("HEX code of %v is %v\n", key, value)
	}
}
