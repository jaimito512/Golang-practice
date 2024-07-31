package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //*testing.T is specifying the type of value that's passing through the function
	//t is the test handler, the value we tell if something goes wrong with test
	d := newDeck()

	if len(d) != 16 { // != returns true if the operands are not equal
		t.Errorf("Expected deck length of 20, but got %v", len(d)) //notify test handler (t) that we got an length different that 20
		//Errorf is recorder
	}

	if d[0] != "Ace of Spades" { //this will check that the first card in deck is Ace of Spades
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" { //this will check that the last card in deck is Four of Clubs
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}


//We write out a few test functions within our Go file.  
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_descktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
