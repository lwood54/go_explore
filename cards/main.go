package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println("toString: ", cards.toString())
	fmt.Println("all the way to byte slice: ", []byte(cards.toString()))
	cards.saveToFile("cardStorage")
}
