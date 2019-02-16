package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	// m rows, n cols
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i][j-1]), float64(dp[i-1][j]))) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[j] = grid[i][j]
				} else {
					dp[j] = dp[j-1] + grid[i][j]
				}
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func main() {
	a := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	fmt.Println(minPathSum(a))
	fmt.Println(minPathSum2(a))
}
