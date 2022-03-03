package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards)
	//printState()

	cards.print()
}

func newCard() string {
	return "asd"
}
