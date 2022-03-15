package main

import "fmt"

func main() {
	// first way
	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#4ba1fi",
	//}

	// second way
	//var colors map[string]string

	//third way
	colors := make(map[string]string)

	//add extra key-value pair
	colors["white"] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colors["white"])

	//delete
	delete(colors, "white")
	fmt.Println(colors)
}
