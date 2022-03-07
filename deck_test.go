package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_deskTesting")

	d := newDeck()
	_ = d.saveToFile("_deskTesting")

	loadedDeck := newDeckFromFile("_deskTesting")
	if len(loadedDeck) != 9 {
		t.Errorf("Expected 9, but got %v", len(d))
	}

	_ = os.Remove("_deskTesting")
}
