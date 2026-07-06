package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tictactoe/internal"
)

func main() {
	game := internal.Game{CurrentPlayer: "X"}
	reader := bufio.NewReader(os.Stdin)
	for {
		game.Board.Print()
		fmt.Printf("Player %s, enter row and col (0 2): ", game.CurrentPlayer)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Parse the string
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Please enter two numbers separated by a space.")
			continue // Skip to next iteration without running game logic
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		game.PlayTurn(x, y)
		winner := game.Board.CheckWinner()
		if winner != "" {
			game.Board.Print()
			fmt.Printf("Player %s wins!\n", winner)
			break // Exit the game loop
		} else if game.Board.IsFull() {
			game.Board.Print()
			fmt.Println("It's a draw!")
			break
		}
	}

}
