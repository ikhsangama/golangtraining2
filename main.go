package main

func main() {
	cards := newDeck()
	hand, remainingCards := cards.deal(5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "asd"
}
