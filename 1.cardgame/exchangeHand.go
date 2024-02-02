package main

import "fmt"

type ExchangeHand struct {
	countDonw int
	exchanger *BasePlayer
	exchangee *BasePlayer
}

func NewExchangeHand(exchanger *BasePlayer, exchnagee *BasePlayer) *ExchangeHand {
	return &ExchangeHand{
		countDonw: 3,
		exchanger: exchanger,
		exchangee: exchnagee,
	}
}

func (e *ExchangeHand) CountDown() {
	e.countDonw--
	if e.countDonw == 0 {
		e.Exechange(true)
	}
}

func (e *ExchangeHand) Exechange(isBack bool) {
	fmt.Printf("exchanger: %s, exchangee: %s isBack: %t\n", e.exchanger.Name, e.exchangee.Name, isBack)
	tmpHand := e.exchangee.Hand
	e.exchangee.Hand = e.exchanger.Hand
	e.exchanger.Hand = tmpHand
}
