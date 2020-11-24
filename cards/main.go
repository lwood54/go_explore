package main

func main() {
	cards := newDeckFromFile("cardStorage")
	cards.shuffle()
	cards.print()
}
