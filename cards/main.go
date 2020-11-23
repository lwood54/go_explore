package main

import "fmt"

func main() {
	// NOTE: Array is fixed length, Slice can have length changed
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Diamonds") // does not mutate original slice, but returns a new slice

	for i, card := range cards {
		// iterating through range, providing index and item
		fmt.Println(i, card)
	}
}

// function that returns value of type string
func newCard() string {
	return "Five of Clubs"
}
