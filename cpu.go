package main()

import (
//"fmt"
// "strconv"
)


func cpu(board [8][8]int, player int) [2]int {
    var zahyoint [2]int
    for j=0;j<=2;j++ { //角のループ
	    if okeru_okenai(j, j, board, player) == true {
	        zahyoint[0] = j
	        zahyoint[1] = j
	        return zahyoint
	    } else if okeru_okenai(j, 7-j, board, player) == true {
	        zahyoint[0] = j
	        zahyoint[1] = 7-j
	        return zahyoint
	    } else if okeru_okenai(7-j, j, board, player) == true {
	        zahyoint[0] = 7-j
	        zahyoint[1] = j
	        return zahyoint
	    } else if okeru_okenai(7-j, 7-j, board, player) == true {
	        zahyoint[0] = 7-j
	        zahyoint[1] = 7-j
	        return zahyoint
	    } else {
	        for i=j; i<=7-j; i++ { //端のループ
	            if okeru_okenai(i, j, board, player) == true {
	                zahyoint[0] = i
	                zahyoint[1] = j
	                return zahyoint
	            }
	        }
	        for i=j; i<=7-j; i++ {
	            if okeru_okenai(j, i, board, player) == true {
	                zahyoint[0] = j
	                zahyoint[1] = i
	                return zahyoint
	            }
	        }
	        for i=j; i<=7-j; i++ {
	            if okeru_okenai(7-j, i, board, player) == true {
	                zahyoint[0] = 7-j
	                zahyoint[1] = i
	                return zahyoint
	            }
	        }
	        for i=j; i<=7-j; i++ {
	            if okeru_okenai(7-j, i, board, player) == true {
	                zahyoint[0] = 7-j
	                zahyoint[1] = i
	                return zahyoint
	            }
	        }
	    }
	}
	return zahyoint
}