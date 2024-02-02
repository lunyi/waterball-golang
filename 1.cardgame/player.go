package main

import (
	"math/rand"
)

type IPlayer interface {
	SelectCard() *Card
	Naming()
	AddHandCard(card *Card)
	AddPoint()
	GetPoint() int
	SetShowdown(showdown *Showdown)
	ShowCards() []*Card
	GetName() string
	GetIndex() int
	GetBasePlayer() *BasePlayer
}

type BasePlayer struct {
	Random       *rand.Rand
	Name         string
	Hand         *Hand
	Point        int
	Index        int
	Showdown     *Showdown
	ExchangeHand *ExchangeHand
}

func NewPlayer(index int) *BasePlayer {
	return &BasePlayer{
		Index:  index,
		Random: rand.New(rand.NewSource(rand.Int63())),
		Point:  0,
		Hand:   &Hand{},
	}
}

func (p *BasePlayer) GetIndex() int {
	return p.Index
}

func (p *BasePlayer) GetName() string {
	return p.Name
}

func (p *BasePlayer) GetBasePlayer() *BasePlayer {
	return p
}

func (p *BasePlayer) ShowCards() []*Card {
	return p.Hand.ShowCards()
}

func (p *BasePlayer) SetShowdown(showdown *Showdown) {
	p.Showdown = showdown
}

func (p *BasePlayer) AddHandCard(card *Card) {
	p.Hand.AddHandCard(card)
}

func (p *BasePlayer) GetPoint() int {
	return p.Point
}

func (p *BasePlayer) AddPoint() {
	p.Point++
}

func GetFinalWinner(players []IPlayer) ([]string, int) {
	var winnerNames []string
	maxPoint := 0
	for _, player := range players {
		if player.GetPoint() > maxPoint {
			maxPoint = player.GetPoint()
			winnerNames = []string{player.GetName()}
		} else if player.GetPoint() == maxPoint {
			winnerNames = append(winnerNames, player.GetName())
		}
	}

	return winnerNames, maxPoint
}
