package main

import "fmt"

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	curProfit, max, buy := 0, 0, prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		if price < buy {
			fmt.Println("Stop curProfit", curProfit, max, " update buy[", i, "]=", price)
			if curProfit > max {
				fmt.Println("update maxProfit to", curProfit)
				max = curProfit
			}
			buy, curProfit = price, 0
		} else {
			diff := price - buy
			fmt.Println("diff", i, price, buy, diff)
			if diff > curProfit {
				fmt.Println("update curProfit", diff)
				curProfit = diff
			}
		}
	}
	if curProfit > max {
		fmt.Println("update maxProfit to", curProfit)
		max = curProfit
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 2, 9, 1, 6, 9}))
}
