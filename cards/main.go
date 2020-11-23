package main

func main() {
	// NOTE: Array is fixed length, Slice can have length changed
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Diamonds") // does not mutate original slice, but returns a new slice

	cards.print()
}

// function that returns value of type string
func newCard() string {
	return "Five of Clubs"
}
