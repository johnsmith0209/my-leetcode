package main

import "fmt"

// [0, 0] -> [0, 3] (row, col) -> (col, len - 1 - row)
// [0, 3] -> [3, 3] (row, col) -> (col, len - 1 - row)
// [3, 3] -> [3, 0] (row, col) -> (col, len - 1 - row)
// [3, 0] -> [0, 0] (row, col) -> (col, len - 1 - row)

// [0, 1] -> [1, 3] (row, col) -> (col, len - 1 - row)
// [1, 3] -> [3, 2] (row, col) -> (col, len - 1 - row)
// [3, 2] -> [2, 0] (row, col) -> (col, len - 1 - row)
// [2, 0] -> [0, 1] (row, col) -> (col, len - 1 - row)

func rotate(matrix [][]int) {
	//roate by circle
	sideLen := len(matrix)
	for c := 0; c < sideLen/2; c++ {
		lineStart, lineEnd := c, sideLen-1-c
		fmt.Println("=============into line ", c, lineStart, lineEnd)
		for i := lineStart; i < lineEnd; i++ {
			fmt.Println("======================into cur step ", c, i)
			// Nth circle i th block, one rotate involves four nodes, four steps
			curNum := matrix[c][i]
			nextRow, nextCol := c, i
			for j := 0; j < 4; j++ {
				nextRow, nextCol = nextCol, sideLen-1-nextRow
				dumpMatrix(matrix)
				fmt.Println("save", curNum, " -> ", nextRow, nextCol)
				curNum, matrix[nextRow][nextCol] = matrix[nextRow][nextCol], curNum
				fmt.Println("after one step rotate")
				dumpMatrix(matrix)
			}
		}
	}
}

func dumpMatrix(matrix [][]int) {
	fmt.Println("-------------------------------------------------")
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println("-------------------------------------------------")
}

func main() {
	// a := [][]int{
	// 	{1, 2, 3, 4, 5, 6},
	// 	{7, 8, 9, 10, 11, 12},
	// 	{13, 14, 15, 16, 17, 18},
	// 	{19, 20, 21, 22, 23, 24},
	// 	{25, 26, 27, 28, 29, 30},
	// 	{31, 32, 33, 34, 35, 36},
	// }
	a := [][]int{
		{1},
	}
	rotate(a)
	dumpMatrix(a)
}

// [1, 1] -> [1, 3] (row, col) -> (col, )
