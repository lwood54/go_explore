package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// (d deck) is a 'receive' of a function
// this establishes that any variable of type 'deck' gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
