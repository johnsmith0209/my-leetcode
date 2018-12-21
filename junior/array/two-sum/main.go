package main

import "fmt"

func twoSum(nums []int, target int) []int {
	rzt, l := []int{}, len(nums)
	for k1, v := range nums {
		if k1 == l-1 {
			return rzt
		}
		for i := k1 + 1; i < l; i++ {
			if v+nums[i] == target {
				rzt = []int{k1, i}
			}
		}
	}
	return rzt
}

func twoSum2(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	for k, v := range nums {
		m[v] = k
	}
	for k, v := range nums {
		find := target - v
		targetKey, ok := m[find]
		if ok && targetKey != k {
			return []int{k, m[find]}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))
}
