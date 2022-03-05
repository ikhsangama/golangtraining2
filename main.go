package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	_ = cards.saveToFile("my_cards")
}

func newCard() string {
	return "asd"
}
