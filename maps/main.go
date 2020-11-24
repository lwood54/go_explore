package main

import "fmt"

func main() {
	// option 1
	// declaring a variable where the keys are of type string and the values are all of type string as well
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// option 2
	// var colors map[string]string

	// option 3
	colors := make(map[string]string)
	// add in key/value pairs after initialization
	colors["white"] = "#ffffff" // cannot access with dot notation, this is because all keys are typed
	fmt.Println(colors)

	// option 3 with ints/strings
	otherColors := make(map[int]string)
	otherColors[10] = "#fff000"
	fmt.Println(otherColors)

	// how to delete a key/value pair from a map
	delete(otherColors, 10)
	fmt.Println(otherColors)

}
