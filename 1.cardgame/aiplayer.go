package main

import (
	"fmt"
	"math/rand"
)

type AIPlayer struct {
	*BasePlayer
}

func NewAIPlayer(index int) IPlayer {
	return &AIPlayer{
		BasePlayer: &BasePlayer{
			Random: rand.New(rand.NewSource(rand.Int63())),
			Point:  0,
			Index:  index,
			Name:   "",
			Hand:   &Hand{},
		},
	}
}

func (a *AIPlayer) SelectCard() *Card {
	if a.ExchangeHand == nil {
		ifExchange := a.Random.Intn(3)
		if ifExchange == 0 {
			exchangees := a.Showdown.GetPlayers()
			for i, player := range exchangees {
				if player.GetIndex() == a.Index {
					exchangees = append(exchangees[:i], exchangees[i+1:]...)
					break
				}
			}

			exchangee := exchangees[a.Random.Intn(len(exchangees))]
			a.ExchangeHand = NewExchangeHand(a.BasePlayer, exchangee.GetBasePlayer())
		}
	} else {
		a.ExchangeHand.CountDown()
	}
	return a.Hand.SelectCard(a.Hand.Size())
}

func (a *AIPlayer) Naming() {
	a.Name = fmt.Sprintf("AI Player-%d", a.Index)
}
