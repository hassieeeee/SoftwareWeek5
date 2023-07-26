package main

import (
//"fmt"
// "strconv"
)

func marking(board [8][8]int, player int) [8][8]int {
	for x := 0; x<=7; x++{
		for y := 0; y<=7; y++{
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
}
