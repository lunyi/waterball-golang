package main

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.Shuffle()
	return deck
}

func (d *Deck) generateCards() {
	for suit := Club; suit <= Spade; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			d.Cards = append(d.Cards, Card{Suit: Suit(suit), Rank: Rank(rank)})
		}
	}
}

func (d *Deck) Shuffle() {
	iterations := 1000
	d.generateCards()
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < iterations; i++ {
		for j := range d.Cards {
			secondCardIndex := rand.Intn(len(d.Cards))
			d.Cards[j], d.Cards[secondCardIndex] = d.Cards[secondCardIndex], d.Cards[j]
		}
	}
}

func (d *Deck) DrawCard() *Card {
	if len(d.Cards) == 0 {
		return nil
	}
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return &card
}

func (d *Deck) getCards() []Card {
	return d.Cards
}
