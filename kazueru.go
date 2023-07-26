func kazueru(board [8][8]int){
    var black int
    var white int
    for i = 0; i < 8; i++ {
	    for j = 0; j < 8; j++ {
			if board[i][j] == 1 {
			    black++
				} 
				else if board[i][j] == 2 {
					white++
				} 
			}
	if(black>white){
	    fmt.Print("black win")
	}
	else if(black < white){
	    fmt.Print("white win")
	}
	else{
	    fmt.Print("draw")
	}
}