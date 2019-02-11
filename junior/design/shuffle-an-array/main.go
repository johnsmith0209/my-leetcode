package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Source, Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{Source: nums, Nums: append([]int{}, nums...)}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.Nums = append([]int{}, this.Source...)
	return this.Nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	// copy source to tmpArr
	l := len(this.Nums)
	for i := 0; i < l; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randIdx := r.Intn(l)
		this.Nums[i], this.Nums[randIdx] = this.Nums[randIdx], this.Nums[i]
	}
	return this.Nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {

	a := Constructor([]int{1, 2, 3, 4, 5, 6})

	fmt.Println("Shuffle", a.Shuffle())
	fmt.Println("Reset", a.Reset())
	fmt.Println("Shuffle", a.Shuffle())
	fmt.Println("Reset", a.Reset())
}
