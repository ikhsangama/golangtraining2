package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected: Ace of Spades, Actual: %v", d[0])
	}

	if d[len(d)-1] != "Three of Diamonds" {
		t.Errorf("Expected: Three of Diamonds, Actual: %v", d[0])
	}

}
