package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Excpected deck lenth of 20 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Excpected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Excpected last card of Five of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadDeckFromFile := newDeckFromFile("_decktesting")
	if len(loadDeckFromFile) != 20 {
		t.Errorf("Excpected deck lenth of 20 but got %v", len(loadDeckFromFile))
	}

	os.Remove("_decktesting")
}
