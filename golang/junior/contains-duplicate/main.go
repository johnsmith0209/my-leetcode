package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] >= 2 {
			return true
		}
	}
	fmt.Println(m)
	return false
}

func containsDuplicate2(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] >= 2 {
			return true
		}
	}
	fmt.Println(m)
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
