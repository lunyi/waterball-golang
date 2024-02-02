package main

type Round struct {
	Card   Card
	Player Player
}

func NewRound(card Card, player Player) *Round {
	return &Round{
		Card:   card,
		Player: player,
	}
}
