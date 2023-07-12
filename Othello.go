package main

import (
	"fmt"
)

// 座標を入力し、x座標、y座標を戻り値
func input() [2]int {
	var zahyostr string
	for {
		fmt.Println("input: ")
		fmt.Scanln(&zahyostr)               // 例：2,1
		x, _ := strconv.Atoi(zahyostr[0:1]) // xにx座標代入
		y, _ := strconv.Atoi(zahyostr[2:3]) // yにy座標代入
		if 1 <= x && x <= 8 && 1 <= y && y <= 8 {
			var zahyoint [2]int = [2]int{x - 1, y - 1}
			break
		}

	}
	return zahyoint
}

func main() {
	for {
		var zahyoint [2]int = input()
		_ = zahyoint

	}
}
