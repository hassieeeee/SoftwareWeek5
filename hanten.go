func hanten (board [8][8]int, player int, x int, y int, direction int) [8][8]int {
    
	if direction == 1 {
		for i:=1;i<7;i++{
		    if board[x-i][y] !=player{
		        board[x-i][y] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 2{
		for i:=1;i<7;i++{
		    if board[x-i][y+i] !=player{
		        board[x-i][y+i] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 3{
		for i:=1;i<7;i++{
		    if board[x][y+i] !=player{
		        board[x][y+i] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 4{
		for i:=1;i<7;i++{
		    if board[x+i][y+i] !=player{
		        board[x+i][y+i] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 5{
		for i:=1;i<7;i++{
		    if board[x+i][y] !=player{
		        board[x+i][y] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 6{
		for i:=1;i<7;i++{
		    if board[x+i][y-i] !=player{
		        board[x+i][y-i] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 7{
		for i:=1;i<7;i++{
		    if board[x][y-i] !=player{
		        board[x][y-i] = player
		    }else{
		    	break
		    }
		}
	}else if direction == 8{
		for i:=1;i<7;i++{
		    if board[x-i][y-i] !=player{
		        board[x-i][y-i] = player
		    }else{
		    	break
		    }
		}
	}
	return board
}





