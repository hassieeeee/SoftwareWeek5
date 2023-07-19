package main

import (
	"fmt"
	//  "reflect"
	"strconv"
)

// func main() {
// 	// 	for i := 0; ; i++{
// 	//      if i%2 == 0 {
// 	//         input()[0]
// 	//      }
// 	//      else{
// 	//          input()[1]
// 	//      }
// 	// 	}
// 	var zahyoint [2]int = input()
// 	_ = zahyoint
// 	fmt.Println(zahyoint[0])
// 	fmt.Println(zahyoint[1])

// }

// 座標を入力し、x座標、y座標を戻り値
func input() [2]int {
	var zahyostr string
	fmt.Println("input: ")
	fmt.Scanln(&zahyostr)               // 例：2,1
	x, _ := strconv.Atoi(zahyostr[0:1]) // xにx座標代入
	y, _ := strconv.Atoi(zahyostr[2:3]) // yにy座標代入
	var zahyoint [2]int = [2]int{x, y}
	return zahyoint
}
