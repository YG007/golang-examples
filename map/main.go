package main

import "fmt"

func main() {

	// first way to assign a value to map
	colors0 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors0)

	// second way to initialize and assign a value to struct
	var colors1 = map[string]string{}
	colors1["white"] = "#ffffff"

	fmt.Println(colors1)

	// third way to initialize and assign a value to struct
	colors2 := make(map[string]string)
	colors2["white"] = "#ffffff"

	fmt.Println(colors2)

	delete(colors2, "white")

	fmt.Println(colors2)

	// iterating over maps
	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap((colors3))
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
