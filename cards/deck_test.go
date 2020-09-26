package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	var cards deck = newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(cards))
	}
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %s", cards[0])
	}
	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %s", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	var cards deck = newDeck()
	cards.saveToFile("_decktesting")
	var loadedDeck deck = newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
