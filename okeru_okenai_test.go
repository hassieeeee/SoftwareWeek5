package test

import "testing"

func TestOkeru_okenai01(t *testing.T) {
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
	var x int = 4
	var y int = 5
	var player int = 1
	hantei := okeru_okenai(x,y,board,player)
	if hantei != true {
		t.Errorf("Test01 is failed")
	}
}