package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// why slow/fast pointer could be faster than this solution?
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		_, ok := m[head]
		if !ok {
			m[head] = true
		} else {
			return true
		}
	}
	return false
}

func main() {

}
