package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// no receiver so tis is a global function, not a method on a deck instance
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) specifies the 'receiver' for the function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(1, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
