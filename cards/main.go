package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// cards := newDeck()
	cards := newDeckFromFile("my_deck.txt")
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// fmt.Println("Remaining Deck:")
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_deck.txt")
}
