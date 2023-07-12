package main

import (
	"fmt"
)

const (
	Empty = 0
	Black = 1
	White = 2
	Size  = 8
)

type Game struct {
	board  [][]int
	player int
}

func NewGame() *Game {
	board := make([][]int, Size)
	for i := range board {
		board[i] = make([]int, Size)
	}

	// initialize board
	board[3][3] = White
	board[3][4] = Black
	board[4][3] = Black
	board[4][4] = White

	return &Game{
		board:  board,
		player: Black,
	}
}
func (g *Game) PrintBoard() {
	fmt.Println("  a b c d e f g h")
	for i := 0; i < Size; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < Size; j++ {
			switch g.board[i][j] {
			case Black:
				fmt.Print("● ")
			case White:
				fmt.Print("○ ")
			default:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
func main() {
	game := NewGame()
	game.PrintBoard()
}
