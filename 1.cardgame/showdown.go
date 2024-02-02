package main

import (
	"fmt"
	"sort"
	"strings"
)

type Showdown struct {
	MapSuit    map[Suit]string
	MapRank    map[Rank]string
	NumOfRanks int
	Deck       *Deck
	Players    []IPlayer
}

func NewShowdown(deck *Deck, players []IPlayer) *Showdown {
	mapSuit := map[Suit]string{
		Heart:   "H",
		Diamond: "D",
		Spade:   "S",
		Club:    "C",
	}

	mapRank := map[Rank]string{
		Ace:   "A",
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
		Ten:   "X",
		Jack:  "J",
		Queen: "Q",
		King:  "K",
	}

	return &Showdown{
		MapSuit:    mapSuit,
		MapRank:    mapRank,
		NumOfRanks: 13,
		Deck:       deck,
		Players:    players,
	}
}

func (s *Showdown) GetPlayers() []IPlayer {
	return s.Players
}

func (s *Showdown) Start() {
	s.initPlayerCards()
	s.runAllRounds()
	s.displayWinner()
}

func (s *Showdown) runAllRounds() {
	for i := 0; i < s.NumOfRanks; i++ {
		fmt.Println("====================")
		fmt.Printf("Round-%d\n", i+1)
		rounds := make([]Round, 0)

		for _, player := range s.Players {
			fmt.Printf("%s\n", player.GetName())
			cards := player.ShowCards()

			for _, card := range cards {
				fmt.Printf(" %s%s ", s.MapRank[card.Rank], s.MapSuit[card.Suit])
			}

			selectedCard := player.SelectCard()

			if selectedCard != nil {
				fmt.Printf(" ==> %s%s\n", s.MapRank[selectedCard.Rank], s.MapSuit[selectedCard.Suit])
				rounds = append(rounds, Round{Player: player, Card: *selectedCard})
			}
		}

		s.compareCard(rounds)
	}
}

func (s *Showdown) compareCard(rounds []Round) {
	sort.Slice(rounds, func(i, j int) bool {
		return rounds[i].Card.CompareTo(rounds[j].Card) == 1
	})

	winningRound := rounds[0]
	winningRound.Player.AddPoint()

	fmt.Printf("round winner: %s %s%s\n", winningRound.Player.GetName(), s.MapRank[winningRound.Card.Rank], s.MapSuit[winningRound.Card.Suit])
}

func (s *Showdown) displayWinner() {
	fmt.Println()
	winners, point := GetFinalWinner(s.Players)

	winnerString := ""
	if len(winners) > 1 {
		winnerString = fmt.Sprintf("Winners are %s", strings.Join(winners, ", "))
	} else {
		winnerString = fmt.Sprintf("The Winner is %s", winners[0])
	}

	result := fmt.Sprintf("(%s and the point is %d)\n", winnerString, point)
	fmt.Println(result)
}

func (s *Showdown) initPlayerCards() {
	for _, player := range s.Players {
		player.Naming()
		player.SetShowdown(s)
	}

	s.Deck.Shuffle()

	for _, player := range s.Players {
		for i := 0; i < s.NumOfRanks; i++ {
			player.AddHandCard(s.Deck.DrawCard())
		}
	}
}
