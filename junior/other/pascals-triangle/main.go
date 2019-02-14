package main

import "fmt"

func generate(numRows int) [][]int {
	rzt := [][]int{}
	for i := 0; i < numRows; i++ {
		if i == 0 {
			rzt = append(rzt, []int{1})
		} else if i == 1 {
			rzt = append(rzt, []int{1, 1})
		} else {
			lastLine, newLine := rzt[len(rzt)-1], []int{1}
			l := len(lastLine)
			for j := 0; j < l; j++ {
				left, right := lastLine[j], 0
				if j != l-1 {
					right = lastLine[j+1]
				}
				newLine = append(newLine, left+right)
			}
			rzt = append(rzt, newLine)
		}
	}
	return rzt
}

func main() {
	fmt.Println(generate(7))
}
