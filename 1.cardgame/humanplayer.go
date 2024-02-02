package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// HumanPlayer type from the provided C# code

type HumanPlayer struct {
	*BasePlayer
}

func NewHumanPlayer(index int) IPlayer {
	return &HumanPlayer{
		BasePlayer: &BasePlayer{
			Random: rand.New(rand.NewSource(rand.Int63())),
			Point:  0,
			Index:  index,
			Name:   "",
			Hand:   &Hand{},
		},
	}
}

func (h *HumanPlayer) SelectCard() *Card {
	if h.ExchangeHand == nil {
		fmt.Printf(" ==> Hi %s, which player do you choose to exchange hand?\n", h.Name)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		playerIndex, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil && playerIndex != h.Index && playerIndex >= 1 && playerIndex <= 4 {
			exchangee := h.Showdown.GetPlayers()[playerIndex-1]
			h.ExchangeHand = NewExchangeHand(h.BasePlayer, exchangee.GetBasePlayer())
		} else {
			fmt.Println("Invalid parameter")
		}
	} else {
		h.ExchangeHand.CountDown()
	}

	return h.Hand.SelectCard(h.Hand.Size())
}

func (h *HumanPlayer) Naming() {
	fmt.Println("Please input your name:")
	reader := bufio.NewReader(os.Stdin)
	inputName, _ := reader.ReadString('\n')
	name := strings.TrimSpace(inputName)

	h.Name = "human"
	if name != "" {
		h.Name = name
	}

	h.Name = fmt.Sprintf("%s-%d", h.Name, h.Index)
}
