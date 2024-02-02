package main

type Round struct {
	Card   Card
	Player IPlayer
}

func NewRound(card Card, player IPlayer) *Round {
	return &Round{
		Card:   card,
		Player: player,
	}
}
