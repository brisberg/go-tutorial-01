package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// (d deck) specifies the 'receiver' for the function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(1, card)
	}
}
