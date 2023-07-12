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

func (g *Game) PlayMove(move string) error { //プレイヤーの移動に応じて盤上に駒を置くことを試みる。移動の正当性が検証され、移動が無効な場合はエラーが返されます。移動が有効であれば、相手の駒を反転させ、別のプレイヤーに切り替えます。
	if len(move) != 2 {
		return fmt.Errorf("Invalid Move, please input a coordinate like a1")
	}

	col := int(move[0] - 'a')
	row := int(move[1] - '1')

	if col < 0 || col >= Size || row < 0 || row >= Size {
		return fmt.Errorf("Invalid Move, out of range")
	}

	if g.board[row][col] != Empty {
		return fmt.Errorf("Invalid Move, coordinate already occupied")
	}

	flipped := g.getFlippedDiscs(row, col)
	if len(flipped) == 0 {
		return fmt.Errorf("Invalid Move, at least 1 disk needs to be flipped")
	}

	g.board[row][col] = g.player
	for _, disc := range flipped {
		g.board[disc[0]][disc[1]] = g.player
	}

	g.player = 3 - g.player // switch player

	return nil
}

func (g *Game) getFlippedDiscs(row, col int) [][]int {
	directions := [][2]int{
		{0, 1},   // 右
		{0, -1},  // 左
		{1, 0},   // 下
		{-1, 0},  // 上
		{1, 1},   // 右下
		{-1, 1},  // 右上
		{1, -1},  // 左下
		{-1, -1}, // 左上
	}

	flipped := [][]int{}
	for _, dir := range directions {
		flipped = append(flipped, g.getFlippedDiscsInDirection(row, col, dir[0], dir[1])...)
	}

	return flipped
}

func (g *Game) getFlippedDiscsInDirection(row, col, dx, dy int) [][]int {
	opp := 3 - g.player
	flipped := [][]int{}
	x, y := row+dx, col+dy

	for x >= 0 && x < Size && y >= 0 && y < Size && g.board[x][y] == opp {
		flipped = append(flipped, []int{x, y})
		x += dx
		y += dy
	}

	if x >= 0 && x < Size && y >= 0 && y < Size && g.board[x][y] == g.player {
		return flipped
	}

	return [][]int{}
}

func main() {
	game := NewGame()
	game.PrintBoard()

	for {
		fmt.Printf("turn of player %s ：", getPlayerName(game.player))
		var move string
		fmt.Scan(&move)

		err := game.PlayMove(move)
		if err != nil {
			fmt.Println("error：", err)
			continue
		}

		game.PrintBoard()

		// 判断
		if !hasValidMove(game) {
			fmt.Println("game over！")
			break
		}
	}
}

func getPlayerName(player int) string {
	switch player {
	case Black:
		return "Black"
	case White:
		return "White"
	default:
		return ""
	}
}

func hasValidMove(g *Game) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if g.board[i][j] == Empty && len(g.getFlippedDiscs(i, j)) > 0 {
				return true
			}
		}
	}
	return false
}
