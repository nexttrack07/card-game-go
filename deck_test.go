package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be Ace of Spades")
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card in deck to be King of Clubs")
	}

	// originalCards := d[:5]
	// d.shuffle()
	// shuffledCards := d[:5]

	// originalCards.print()
	// shuffledCards.print()

	// for i, card := range originalCards {
	// 	if shuffledCards[i] == card {
	// 		t.Error("Shuffle not working")
	// 	}
	// // }

	// d[:5].print()
	// d.shuffle()
	// d[:5].print()

	hand, remainingCards := deal(d, 5)

	if len(hand) != 5 {
		t.Error("deal not working")
	} else if len(remainingCards) != 47 {
		t.Error("deal not working")
	} else if len(d) != 52 {
		t.Error("deal mutates deck!")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
