package internal

import "fmt"

type Board struct {
	Cells [3][3]string
}

func (b *Board) Print() {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == "" {
				fmt.Print("[ ]")
			} else {
				fmt.Printf("[%s]", cell)
			}
		}
		fmt.Println()
	}
}

func (b *Board) Place(x, y int, symbol string) bool {
	if x < 0 || x > 2 || y < 0 || y > 2 || b.Cells[x][y] != "" {
		return false
	}
	b.Cells[x][y] = symbol
	return true
}

func (b *Board) CheckWinner() string {
	for i := 0; i < 3; i++ {
		// Rows
		if b.Cells[i][0] != "" && b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][1] == b.Cells[i][2] {
			return b.Cells[i][0]
		}
		// Columns
		if b.Cells[0][i] != "" && b.Cells[0][i] == b.Cells[1][i] && b.Cells[1][i] == b.Cells[2][i] {
			return b.Cells[0][i]
		}
	}

	if b.Cells[0][0] != "" && b.Cells[0][0] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][2] {
		return b.Cells[0][0]
	}
	if b.Cells[0][2] != "" && b.Cells[0][2] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][0] {
		return b.Cells[0][2]
	}

	return ""
}

func (b *Board) IsFull() bool {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}
