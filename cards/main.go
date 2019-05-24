package main

import "fmt"

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("my_deck.txt")

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println("Remaining Deck:")
	// remainingDeck.print()

	fmt.Println(cards.toString())
	cards.saveToFile("my_deck.txt")
}
