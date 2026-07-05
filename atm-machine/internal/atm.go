package internal

type IATM interface {
	SetState(state State)
	InsertCard(card Card)
	EnterPin(pin int)
	Withdraw(amount int)
	EjectCard()
}

type ATM struct {
	currentState State
}

// EjectCard implements [IATM].
func (a *ATM) EjectCard() {
	a.currentState.EjectCard(a)
}

// EnterPin implements [IATM].
func (a *ATM) EnterPin(pin int) {
	a.currentState.EnterPin(a, pin)
}

// InsertCard implements [IATM].
func (a *ATM) InsertCard(card Card) {
	a.currentState.InsertCard(a, card)
}

// Withdraw implements [IATM].
func (a *ATM) Withdraw(amount int) {
	a.currentState.Withdraw(a, amount)
}

func NewATM(state State) IATM {
	return &ATM{currentState: state}
}

func (a *ATM) SetState(state State) {
	a.currentState = state
}
