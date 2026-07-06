package internal

import "fmt"

type Game struct {
	Board         Board
	CurrentPlayer string
}

func (g *Game) PlayTurn(x, y int) {
	if g.Board.Place(x, y, g.CurrentPlayer) {
		g.CurrentPlayer = map[string]string{"X": "O", "O": "X"}[g.CurrentPlayer]
	} else {
		fmt.Println("Invalid move, try again!")
	}
}
