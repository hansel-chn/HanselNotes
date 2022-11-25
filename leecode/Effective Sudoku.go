package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func isValidSudoku(board [][]byte) bool {
	var rows [9][9]int
	var cols [9][9]int
	// go定义切片从后往前想
	var blocks [3][3][9]int

	for row, data := range board {
		for col, val := range data {
			if val == '.' {
				continue
			}
			val := val - '1'
			rows[row][val]++
			cols[col][val]++
			blocks[row/3][col/3][val]++
			if rows[row][val] > 1 || cols[col][val] > 1 || blocks[row/3][col/3][val] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	var board = [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}
	//var transfer [][]byte
	//var transfer = [][]byte{}
	var transfer = make([][]byte, 9)

	for i, data := range board {

		fmt.Println(data[0])
		//fmt.Println([]byte(data))
		//fmt.Println(byte(data[0]))
		fmt.Println(data[0][0])
		fmt.Println(byte(data[0][0]))
		fmt.Println(byte(data[1][0]))
		fmt.Println(reflect.TypeOf(byte(data[0][0])))

		fmt.Println(reflect.TypeOf(data))
		tempData := strings.Join(data, "")
		transfer[i] = []byte(tempData)
		fmt.Println(transfer[i])
		fmt.Printf("%s", string(transfer[i]))
		os.Exit(1)
		//transfer = append(transfer, []byte(data[0]))

		//fmt.Println(reflect.TypeOf([]byte(tempData)))
		//transfer[i] = []byte(tempData)

		//fmt.Println([]byte(tempData))
		//fmt.Println(tempData)
	}
	//for i, data := range board {
	//	for j, val := range data {
	//		transfer[i][j] = val[0]
	//		fmt.Println(transfer[i][j])
	//		//fmt.Println([]byte(tempData))
	//		//fmt.Println(tempData)
	//	}
	//}
	judge := isValidSudoku(transfer)
	fmt.Println(judge)
}
