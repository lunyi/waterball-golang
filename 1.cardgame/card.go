package main

type Suit int

const (
	Club Suit = iota + 1
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	return [...]string{"C", "D", "H", "S"}[s-1]
}

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r-2]
}

type Card struct {
	Suit Suit
	Rank Rank
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

func (c Card) CompareTo(card Card) int {
	rankCompare := c.Rank - card.Rank
	if rankCompare != 0 {
		return int(rankCompare)
	}
	return int(c.Suit - card.Suit)
}
