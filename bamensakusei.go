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
func (g *Game) PrintBoard() {
	fmt.Println("  1 2 3 4 5 6 7 8")
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
