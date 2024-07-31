package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new custom type, of 'deck'
// which is a slice of strings
type deck []string  

func newDeck() deck { //anytime a someone calls for the newDeck function, this function will create a new deck
						//this function returns a value of type 'deck', which is cards.
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"} //String Slice of card Suits
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
func (d deck) print() { //here we created a function that has a 'receiver' before the function name, print.
						// the Receiver here is "d", of type 'deck'
						// the function is 'print'
	for i, card := range d {
		fmt.Println(i, card)
	}

	//d.print() //we are able to call a method of print on "d"
}


// Here's a function, deal, without any receiver.  We choose to take in a deck (d) as an argument
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //we were able to create slices (arrays) of different contents
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") //Ace of Spades,Two of Spades, Three of Spades...
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
