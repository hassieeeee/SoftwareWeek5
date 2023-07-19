package main

import (
	"fmt"
	"strconv"
)

const (
	Empty = 0
	Black = 1
	White = 2
	Size  = 8
)

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
	PrintBoard(board)

	player := Black
	fmt.Println("player:%d please input x,y\n", player)
	var zahyoint [2]int = input()
	var okeru bool = okeru_okenai(zahyoint[0], zahyoint[1], board, player)
	if okeru == true {
		board = komahaitti(zahyoint[0], zahyoint[1], board, player)
	} else {
		println("置けません")
	}
	PrintBoard(board)
	//else
	// もう一回入力を求める
	//配置されたコマを基に白黒反転
	//PrintBoard()

	//勝敗を判断
	//手番交代 player = player % 2 + 1

	// for {
	// 	var zahyoint [2]int = input()
	// 	_ = zahyoint

	// }
}

//座標を入力し、x座標、y座標を戻り値
func input() [2]int {
	var zahyostr string
	var zahyoint [2]int
	for {
		fmt.Println("input: ")
		fmt.Scanln(&zahyostr)               // 例：2,1
		x, _ := strconv.Atoi(zahyostr[0:1]) // xにx座標代入
		y, _ := strconv.Atoi(zahyostr[2:3]) // yにy座標代入
		if 1 <= x && x <= 8 && 1 <= y && y <= 8 {
			zahyoint = [2]int{x - 1, y - 1}
			break
		}
	}
	return zahyoint
}

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

func okeru_okenai(x int, y int, board [8][8]int, player int) bool {
	if board[y][x] != 0 {
		return false
	} else {
		if Top(x, y, board, player) == true {
			return true
		}
		if TopRight(x, y, board, player) == true {
			return true
		}
		if Right(x, y, board, player) == true {
			return true
		}
		if BottomRight(x, y, board, player) == true {
			return true
		}
		if Bottom(x, y, board, player) == true {
			return true
		}
		if BottomLeft(x, y, board, player) == true {
			return true
		}
		if Left(x, y, board, player) == true {
			return true
		}
		if TopLeft(x, y, board, player) == true {
			return true
		}
	}
	return false
}

func Top(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 0 || y == 1 {
		return hantei
	} else {
		if board[y-1][x] != player%2+1 {
			return hantei
		} else {
			for i := 2; y-i >= 0; i++ {
				if board[y-i][x] == player {
					hantei = true
					return hantei
				} else if board[y-i][x] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func TopRight(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 0 || y == 1 || x == 6 || x == 7 {
		return hantei
	} else {
		if board[y-1][x+1] != player%2+1 {
			return hantei
		} else {
			for i := 2; y-i >= 0 && x+i <= 7; i++ {
				if board[y-i][x+i] == player {
					hantei = true
					return hantei
				} else if board[y-i][x+i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func Right(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if x == 6 || x == 7 {
		return hantei
	} else {
		if board[y][x+1] != player%2+1 {
			return hantei
		} else {
			for i := 2; x+i <= 7; i++ {
				if board[y][x+i] == player {
					hantei = true
					break
				} else if board[y][x+i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}
func BottomRight(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 6 || y == 7 || x == 6 || x == 7 {
		return hantei
	} else {
		if board[y+1][x+1] != player%2+1 {
			return hantei
		} else {
			for i := 2; y+i <= 7 && x+i <= 7; i++ {
				if board[y+i][x+i] == player {
					hantei = true
					return hantei
				} else if board[y+i][x+i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func Bottom(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 6 || y == 7 {
		return hantei
	} else {
		if board[y+1][x] != player%2+1 {
			return hantei
		} else {
			for i := 2; y+i <= 7; i++ {
				if board[y+i][x] == player {
					hantei = true
					break
				} else if board[y+i][x] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func BottomLeft(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 6 || y == 7 || x == 0 || x == 1 {
		return hantei
	} else {
		if board[y+1][x-1] != player%2+1 {
			return hantei
		} else {
			for i := 2; y+i <= 7 && x-i >= 0; i++ {
				if board[y+i][x-i] == player {
					hantei = true
					return hantei
				} else if board[y+i][x-i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func Left(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if x == 0 || x == 1 {
		return hantei
	} else {
		if board[y][x-1] != player%2+1 {
			return hantei
		} else {
			for i := 2; x-i >= 0; i++ {
				if board[y][x-i] == player {
					hantei = true
					break
				} else if board[y][x-i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func TopLeft(x int, y int, board [8][8]int, player int) bool {
	hantei := false
	if y == 0 || y == 1 || x == 0 || x == 1 {
		return hantei
	} else {
		if board[y-1][x-1] != player%2+1 {
			return hantei
		} else {
			for i := 2; y-i >= 0 && x-i >= 0; i++ {
				if board[y-i][x-i] == player {
					hantei = true
					return hantei
				} else if board[y-i][x-i] == 0 {
					break
				}
			}
		}
	}
	return hantei
}

func komahaitti(x int, y int, board [8][8]int, player int) [8][8]int {
	board[y][x] = player
	return board
}
