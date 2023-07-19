package main

import (
	"fmt"
	// "strconv"
)

const (
	Empty = 0
	Black = 1
	White = 2
	Size  = 8
)

// 座標を入力し、x座標、y座標を戻り値
// func input() [2]int {
// 	var zahyostr string
// 	for {
// 		fmt.Println("input: ")
// 		fmt.Scanln(&zahyostr)               // 例：2,1
// 		x, _ := strconv.Atoi(zahyostr[0:1]) // xにx座標代入
// 		y, _ := strconv.Atoi(zahyostr[2:3]) // yにy座標代入
// 		if 1 <= x && x <= 8 && 1 <= y && y <= 8 {
// 			var zahyoint [2]int = [2]int{x - 1, y - 1}
// 			break
// 		}

// 	}
// 	return zahyoint
// }

func PrintBoard(board [8][8]int) {
	fmt.Println("  1 2 3 4 5 6 7 8")
	for i := 0; i < Size; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < Size; j++ {
			switch board[i][j] {
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
	var board = [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	//var player int
	PrintBoard(board)

	// for {
	// 	var zahyoint [2]int = input()
	// 	_ = zahyoint

	// }
}
