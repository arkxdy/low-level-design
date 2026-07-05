package internal

import "fmt"

type State interface {
	InsertCard(atm *ATM, card Card)
	EnterPin(atm *ATM, pin int)
	Withdraw(atm *ATM, amount int)
	EjectCard(atm *ATM)
}

type Card struct {
	CardNumber int
	CardName   string
}

// --- IdleState ---
type IdleState struct{}

func (s *IdleState) InsertCard(atm *ATM, card Card) {
	fmt.Println("Card inserted.")
	atm.SetState(&CardInsertedState{})
}
func (s *IdleState) EnterPin(atm *ATM, pin int)    { fmt.Println("No card.") }
func (s *IdleState) Withdraw(atm *ATM, amount int) { fmt.Println("No card.") }
func (s *IdleState) EjectCard(atm *ATM)            { fmt.Println("No card.") }

// --- CardInsertedState ---
type CardInsertedState struct{}

func (s *CardInsertedState) InsertCard(atm *ATM, card Card) { fmt.Println("Already inserted.") }
func (s *CardInsertedState) EnterPin(atm *ATM, pin int) {
	if pin == 1234 {
		fmt.Println("PIN correct.")
		atm.SetState(&AuthenticatedState{})
	} else {
		fmt.Println("Wrong PIN.")
	}
}
func (s *CardInsertedState) Withdraw(atm *ATM, amount int) { fmt.Println("Enter PIN first.") }
func (s *CardInsertedState) EjectCard(atm *ATM) {
	fmt.Println("Card ejected.")
	atm.SetState(&IdleState{})
}

// --- AuthenticatedState ---
type AuthenticatedState struct{}

func (s *AuthenticatedState) InsertCard(atm *ATM, card Card) { fmt.Println("Already authenticated.") }
func (s *AuthenticatedState) EnterPin(atm *ATM, pin int)     { fmt.Println("Already authenticated.") }
func (s *AuthenticatedState) Withdraw(atm *ATM, amount int) {
	fmt.Printf("Withdrawing %d...\n", amount)
	atm.SetState(&TransactionProcessingState{})
}
func (s *AuthenticatedState) EjectCard(atm *ATM) {
	fmt.Println("Card ejected.")
	atm.SetState(&IdleState{})
}

// --- TransactionProcessingState ---
type TransactionProcessingState struct{}

func (s *TransactionProcessingState) InsertCard(atm *ATM, card Card) { fmt.Println("Processing...") }
func (s *TransactionProcessingState) EnterPin(atm *ATM, pin int)     { fmt.Println("Processing...") }
func (s *TransactionProcessingState) Withdraw(atm *ATM, amount int)  { fmt.Println("Processing...") }
func (s *TransactionProcessingState) EjectCard(atm *ATM) {
	fmt.Println("Cannot eject while processing.")
}
