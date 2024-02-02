package main

type Hand struct {
	Cards []*Card
}

func (hand *Hand) Size() int {
	return len(hand.Cards)
}

func (hand *Hand) SelectCard(index int) *Card {
	if index == 0 {
		return nil
	}
	card := hand.Cards[index]
	hand.Cards = hand.Cards[:len(hand.Cards)-1]
	return card
}

func (hand *Hand) ShowCards() []*Card {
	return hand.Cards
}

func (hand *Hand) AddHandCard(card *Card) {
	hand.Cards = append(hand.Cards, card)
}
