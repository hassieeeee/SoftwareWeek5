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