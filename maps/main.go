package main

import "fmt"

func main() {
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors["white"] = "#ffffff"

	fmt.Println(colors)

	delete(colors, "white") // delete this key value
}
