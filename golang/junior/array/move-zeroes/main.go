package main

import "fmt"

func moveZeroes(nums []int) {
	leadingIdx, lastIdx, l := 0, 0, len(nums)
	for leadingIdx < l {
		n := nums[leadingIdx]
		if n == 0 {
			leadingIdx++
		} else {
			fmt.Println("before switch", nums, lastIdx, leadingIdx)
			nums[lastIdx], nums[leadingIdx] = nums[leadingIdx], nums[lastIdx]
			leadingIdx++
			lastIdx++
			fmt.Println("after switch", nums, lastIdx, leadingIdx)
		}
	}
}

func moveZeroes2(nums []int) {
	n, i := len(nums), 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for ; i < n; i++ {
		nums[i] = 0
	}
}

func main() {
	a := []int{0, 0, 1, 0, 0, 2, 0, 3, 0, 4, 5, 6}
	moveZeroes2(a)
	fmt.Println(a)
}
