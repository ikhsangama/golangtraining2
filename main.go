package main

import "fmt"

func main() {
	cards := deck{newCard(), "2"}
	cards = append(cards, "6")
	fmt.Println(cards)
	//printState()

	cards.print()
}

func newCard() string {
	return "asd"
}
