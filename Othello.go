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

	player := Black
	skip := 0

	for {
		fmt.Println("Please select a mode. vsCPU: c vsPlayer: p")
		var mode string
		fmt.Scanln(&mode)
		if mode == "c" { //CPU対戦
			for {
				board = marking(board, player)
				PrintBoard(board)
				if skipjudge(board) == false { //おける場所が無かったらplayerを変更してスキップ
					skip = skip + 1
					player = player%2 + 1
					fmt.Println("skip!")
					if skip == 2 {
						break //両方連続でスキップしたら置く場所が無いのでゲーム終了して勝敗判定へ
					}
				} else {
					skip = 0
					board = deletemark(board)
					if player == Black {
						fmt.Println("player: ● please input x,y")
					} else {
						fmt.Println("player: ○ please input x,y")
					}
					var zahyoint [2]int
					if player == Black {
						zahyoint = input()
					} else if player == White {
						zahyoint = cpu(board, player)
					}
					var okeru bool = okeru_okenai(zahyoint[0], zahyoint[1], board, player)
					if okeru == true {
						board = komahaitti(zahyoint[0], zahyoint[1], board, player)
						board = hasamu(zahyoint[0], zahyoint[1], board, player)
						player = player%2 + 1
					} else {
						println("Cannot place. Please try again.")
					}
				}
			}
			kazueru(board) //勝敗を判断
			break

		} else if mode == "p" { //2P対戦
			for {
				board = marking(board, player)
				PrintBoard(board)
				if skipjudge(board) == false { //おける場所が無かったらplayerを変更してスキップ
					skip = skip + 1
					player = player%2 + 1
					if skip == 2 {
						break //両方連続でスキップしたら置く場所が無いのでゲーム終了して勝敗判定へ
					}
				} else {
					skip = 0
					board = deletemark(board)
					if player == Black {
						fmt.Println("player: ● please input x,y")
					} else {
						fmt.Println("player: ○ please input x,y")
					}
					var zahyoint [2]int = input()
					var okeru bool = okeru_okenai(zahyoint[0], zahyoint[1], board, player)
					if okeru == true {
						board = komahaitti(zahyoint[0], zahyoint[1], board, player)
						board = hasamu(zahyoint[0], zahyoint[1], board, player)
						player = player%2 + 1
					} else {
						println("Cannot place. Please try again.")
					}
				}
			}
			kazueru(board) //勝敗を判断
			break
		} else {
			fmt.Println("Invalid input.")
		}
	}

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

//現在の盤面を返す
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
			case 3:
				fmt.Print("x ")
			default:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

//おける座標にマークをつけて盤面を返す
func marking(board [8][8]int, player int) [8][8]int {
	for x := 0; x <= 7; x++ {
		for y := 0; y <= 7; y++ {
			if board[y][x] == 0 {
				if Top(x, y, board, player) == true {
					board[y][x] = 3
				}
				if TopRight(x, y, board, player) == true {
					board[y][x] = 3
				}
				if Right(x, y, board, player) == true {
					board[y][x] = 3
				}
				if BottomRight(x, y, board, player) == true {
					board[y][x] = 3
				}
				if Bottom(x, y, board, player) == true {
					board[y][x] = 3
				}
				if BottomLeft(x, y, board, player) == true {
					board[y][x] = 3
				}
				if Left(x, y, board, player) == true {
					board[y][x] = 3
				}
				if TopLeft(x, y, board, player) == true {
					board[y][x] = 3
				}
			}
		}
	}
	return board
}

//マークを消した盤面を返す
func deletemark(board [8][8]int) [8][8]int {
	for x := 0; x <= 7; x++ {
		for y := 0; y <= 7; y++ {
			if board[y][x] == 3 {
				board[y][x] = 0
			}
		}
	}
	return board
}

//指定した座標におけるかどうかをbool型で返す
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

func hasamu(x int, y int, board [8][8]int, player int) [8][8]int {
	if Top(x, y, board, player) == true {
		board = hanten(board, player, x, y, 1)
	}
	if TopRight(x, y, board, player) == true {
		board = hanten(board, player, x, y, 2)
	}
	if Right(x, y, board, player) == true {
		board = hanten(board, player, x, y, 3)
	}
	if BottomRight(x, y, board, player) == true {
		board = hanten(board, player, x, y, 4)
	}
	if Bottom(x, y, board, player) == true {
		board = hanten(board, player, x, y, 5)
	}
	if BottomLeft(x, y, board, player) == true {
		board = hanten(board, player, x, y, 6)
	}
	if Left(x, y, board, player) == true {
		board = hanten(board, player, x, y, 7)
	}
	if TopLeft(x, y, board, player) == true {
		board = hanten(board, player, x, y, 8)
	}
	return board
}

func hanten(board [8][8]int, player int, y int, x int, direction int) [8][8]int {

	if direction == 1 {
		for i := 1; x-i >= 0; i++ {
			if board[x-i][y] == player%2+1 {
				board[x-i][y] = player
			} else {
				break
			}
		}
	} else if direction == 2 {
		for i := 1; i < 7; i++ {
			if board[x-i][y+i] == player%2+1 {
				board[x-i][y+i] = player
			} else {
				break
			}
		}
	} else if direction == 3 {
		for i := 1; i < 7; i++ {
			if board[x][y+i] == player%2+1 {
				board[x][y+i] = player
			} else {
				break
			}
		}
	} else if direction == 4 {
		for i := 1; i < 7; i++ {
			if board[x+i][y+i] == player%2+1 {
				board[x+i][y+i] = player
			} else {
				break
			}
		}
	} else if direction == 5 {
		for i := 1; i < 7; i++ {
			if board[x+i][y] == player%2+1 {
				board[x+i][y] = player
			} else {
				break
			}
		}
	} else if direction == 6 {
		for i := 1; i < 7; i++ {
			if board[x+i][y-i] == player%2+1 {
				board[x+i][y-i] = player
			} else {
				break
			}
		}
	} else if direction == 7 {
		for i := 1; y-i >= 0; i++ {
			if board[x][y-i] == player%2+1 {
				board[x][y-i] = player
			} else {
				break
			}
		}
	} else if direction == 8 {
		for i := 1; i < 7; i++ {
			if board[x-i][y-i] == player%2+1 {
				board[x-i][y-i] = player
			} else {
				break
			}
		}
	}
	return board
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

func kazueru(board [8][8]int) {
	var black int
	var white int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 1 {
				black++
			} else if board[i][j] == 2 {
				white++
			}
		}
	}
	if black > white {
		fmt.Print("black win")
	} else if black < white {
		fmt.Print("white win")
	} else {
		fmt.Print("draw")
	}
}

func skipjudge(board [8][8]int) bool {
	for x := 0; x <= 7; x++ {
		for y := 0; y <= 7; y++ {
			if board[y][x] == 3 {
				return true
			}
		}
	}
	return false
}

func cpu(board [8][8]int, player int) [2]int {
	var zahyoint [2]int
	for j := 0; j <= 2; j++ { //角のループ
		if okeru_okenai(j, j, board, player) == true {
			zahyoint[0] = j
			zahyoint[1] = j
			return zahyoint
		} else if okeru_okenai(j, 7-j, board, player) == true {
			zahyoint[0] = j
			zahyoint[1] = 7 - j
			return zahyoint
		} else if okeru_okenai(7-j, j, board, player) == true {
			zahyoint[0] = 7 - j
			zahyoint[1] = j
			return zahyoint
		} else if okeru_okenai(7-j, 7-j, board, player) == true {
			zahyoint[0] = 7 - j
			zahyoint[1] = 7 - j
			return zahyoint
		} else {
			for i := j; i <= 7-j; i++ { //端のループ
				if okeru_okenai(i, j, board, player) == true {
					zahyoint[0] = i
					zahyoint[1] = j
					return zahyoint
				}
			}
			for i := j; i <= 7-j; i++ {
				if okeru_okenai(j, i, board, player) == true {
					zahyoint[0] = j
					zahyoint[1] = i
					return zahyoint
				}
			}
			for i := j; i <= 7-j; i++ {
				if okeru_okenai(7-j, i, board, player) == true {
					zahyoint[0] = 7 - j
					zahyoint[1] = i
					return zahyoint
				}
			}
			for i := j; i <= 7-j; i++ {
				if okeru_okenai(7-j, i, board, player) == true {
					zahyoint[0] = 7 - j
					zahyoint[1] = i
					return zahyoint
				}
			}
		}
	}
	return zahyoint
}
