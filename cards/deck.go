package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) is a 'receiver' of a function
// this establishes that any variable of type 'deck' gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function takes 2 specific parameters, type deck and type int, returns to specific values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// a method that allows a receiver of type deck to save data to a file
// accepts a parameter of filename type string and returns an error (which is what WriteFile returns)
func (d deck) saveToFile(filename string) error {
	// return error generated from WriteFile, 3 parameters passed
	// conversion occuring in 2nd param, and write permissions of everyone for 3rd param
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
