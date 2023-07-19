package main

//import "fmt"

// const (
// 	Empty = 0
// 	Black = 1
// 	White = 2
// 	Size  = 8
// )

// func PrintBoard(board [8][8]int) {
// 	fmt.Println("  1 2 3 4 5 6 7 8")
// 	for i := 0; i < Size; i++ {
// 		fmt.Printf("%d ", i+1)
// 		for j := 0; j < Size; j++ {
// 			switch board[i][j] {
// 			case Black:
// 				fmt.Print("● ")
// 			case White:
// 				fmt.Print("○ ")
// 			default:
// 				fmt.Print(". ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	var x, y int
// 	var n = 1
// 	var board = [8][8]int{
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 0, 0, 2, 1, 0, 0, 0},
// 		{0, 0, 0, 1, 2, 0, 0, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 		{0, 0, 0, 0, 0, 0, 0, 0},
// 	}
// 	PrintBoard(board)

// 	var player = Black

// 	for {

// 		for {
// 			fmt.Println("player: please input x,y\n", player)
// 			fmt.Println(player)
// 			fmt.Print("x = ")
// 			fmt.Scanln(&y)
// 			fmt.Print("y = ")
// 			fmt.Scanln(&x)
// 			x--
// 			y--
// 			if 0 <= x && x <= 7 && 0 <= y && y <= 7 && board[x][y] == 0 {
// 				board[x][y] = player
// 				break
// 			} //else if x == -1 && y == -1 {
// 			//player = player%2 + 1
// 			//} else if x == 11 && y == 11 {
// 			//break
// 			//}

// 		}
// 		player = n%2 + 1
// 		PrintBoard(board)

// 	}

// }

func komahaitti(x int, y int, board [8][8]int, player int) [8][8]int {
	board[y][x] = player
	return board
}
