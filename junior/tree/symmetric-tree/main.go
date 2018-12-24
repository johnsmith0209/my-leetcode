package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	return isMirror(root, root)
}

func isSymmetric(root *TreeNode) bool {
	childCount, children, level, mark := 0, []*TreeNode{}, 0, -12345678
	if root == nil {
		return true
	}
	childCount, children = 1, append(children, root)
	for childCount > 0 {
		level++
		nums, newChildren, blank := make([]int, len(children)), []*TreeNode{}, true
		for k, v := range children {
			if v == nil {
				fmt.Println("v is nil, nums skip")
				nums[k] = mark
				continue
			} else {
				fmt.Printf("v[%d] into nums \n", v.Val)
				nums[k] = v.Val
			}
			if v.Left != nil {
				fmt.Printf("v.left[%d] into children \n", v.Left.Val)
				blank, newChildren = false, append(newChildren, v.Left)
			} else {
				fmt.Println("v.left nil into children")
				newChildren = append(newChildren, nil)
			}

			if v.Right != nil {
				fmt.Printf("v.right[%d] into children \n", v.Right.Val)
				blank, newChildren = false, append(newChildren, v.Right)
			} else {
				fmt.Println("v.right nil into children")
				newChildren = append(newChildren, nil)
			}
		}
		fmt.Println("level", level, "nums", nums, "newChildren", len(newChildren))
		if len(nums) != len(children) {
			fmt.Println("level", level, "nums", nums, "not equal to children, return false")
			return false
		}
		l := len(nums)
		a, b := nums[0:l/2], nums[l/2:]
		for i := 0; i < l/2; i++ {
			if a[i] != b[l/2-1-i] {
				fmt.Printf("a[%d]=%d, b[%d]=%d, not equal", i, a[i], l/2-1-i, b[l/2-1-i])
				return false
			}
		}
		fmt.Println("nums split into ", a, b)
		if blank {
			break
		}
		children, childCount = newChildren, len(newChildren)
	}
	return true
}

func main() {

}
