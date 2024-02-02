package main

func main() {
	deck := NewDeck()
	players := []IPlayer{
		NewHumanPlayer(1),
		NewAIPlayer(2),
		NewAIPlayer(3),
		NewAIPlayer(4),
	}
	showdown := NewShowdown(deck, players)
	showdown.Start()
}
