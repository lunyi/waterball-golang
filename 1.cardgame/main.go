package main

import "fmt"

func main() {
	deck := NewDeck()
	cards := deck.getCards()

	for i := 0; i < len(cards); i++ {
		fmt.Printf("%s[%s],", cards[i].Rank, cards[i].Suit)
	}
}
