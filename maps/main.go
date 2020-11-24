package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

//// MAPS  vs STRUCTS ////
//					Map															//														Struct						//
//	1. All values must be the same type			//	1. Values can be of different type
//	2. Use to represent a collection				//	2. Used to represent a 'thing' with
//		of related properties													a lot of different properties
//	3. Keys are indexed - we can iterate		//	3. Keys don't support indexing
//			over them
//	4. Don't need to know all the keys			//	4. You need to know all the different fields at compile time
//			at compile time
//	5. Reference Type												//	5. Value Type
//	6. All keys must be the same type
