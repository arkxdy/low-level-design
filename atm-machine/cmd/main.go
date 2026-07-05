package cmd

import "atm-machine/internal"

func main() {
	atm := internal.NewATM(&internal.IdleState{})

	card := &internal.Card{
		CardNumber: 1234,
		CardName:   "John",
	}

	atm.InsertCard(*card)
	atm.EnterPin(1234)
	atm.Withdraw(12000)
	atm.EjectCard()
}
