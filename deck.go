package main

import "fmt"

// Create a new type of 'deck', which is a slice of string
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	var cards = deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		card := suit
		fmt.Println(card, "<<<1")
		for _, value := range cardValues {
			fmt.Println(card, "<<<")
			card = value + " of " + card
			cards = append(cards, card)
		}
	}

	return cards
}
