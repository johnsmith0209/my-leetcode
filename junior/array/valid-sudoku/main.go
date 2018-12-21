package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	fmt.Println(board)
	// row
	for i := 0; i < 9; i++ {
		if findDuplicate(extractNums(board, "row", i)) {
			fmt.Println("find row duplicate", i)
			return false
		}
	}
	// cols
	for i := 0; i < 9; i++ {
		if findDuplicate(extractNums(board, "col", i)) {
			fmt.Println("find col duplicate", i)
			return false
		}
	}
	// cube
	for i := 0; i < 9; i++ {
		if findDuplicate(extractNums(board, "cube", i)) {
			fmt.Println("find col duplicate", i)
			return false
		}
	}
	return true
}

func findDuplicate(nums []byte) bool {
	m := make(map[byte]int, len(nums))
	for k, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		if v != '.' {
			m[v] = k
		}
	}
	return false
}

func extractNums(board [][]byte, numType string, start int) []byte {
	rzt := []byte{}
	switch numType {
	case "row":
		fmt.Println("get row", start)
		rzt = board[start]
	case "col":
		fmt.Println("get col", start)
		for i := 0; i < 9; i++ {
			rzt = append(rzt, board[i][start])
		}
	case "cube":
		// row start/end, col start/end
		mod, multi := start%3, start/3
		rS, rE, cS, cE := 3*multi, 3*multi+2, 3*mod, 3*mod+2
		fmt.Println("get cube", start, rS, rE, cS, cE)
		for i := rS; i <= rE; i++ {
			for j := cS; j <= cE; j++ {
				rzt = append(rzt, board[i][j])
			}
		}
	}
	return rzt
}

func main() {
	a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(a))
}
