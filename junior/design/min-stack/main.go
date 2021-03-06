package main

import "fmt"

type MinStack struct {
	stack, minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = []int{x}
		this.minStack = []int{x}
		return
	}
	min := this.GetMin()
	// multiple push when equals to min
	if x <= min {
		this.minStack = append(this.minStack, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	last, min := this.stack[len(this.stack)-1], this.GetMin()
	if min == last {
		this.minStack = this.minStack[0 : len(this.minStack)-1]
	}
	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	stack := Constructor()
	// stack.Pop()
	stack.Push(10)
	stack.Push(1)
	stack.Push(10)
	stack.Push(1)
	stack.Push(1)
	fmt.Println("!!!! get min", stack.GetMin(), "get top", stack.Top())
	stack.Pop()
	fmt.Println("!!!! get min", stack.GetMin(), "get top", stack.Top())
	stack.Pop()
	stack.Pop()
	stack.Pop()
	fmt.Println("!!!! get min", stack.GetMin(), "get top", stack.Top())
}
