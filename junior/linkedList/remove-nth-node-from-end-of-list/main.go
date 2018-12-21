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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	var s []*ListNode
	for true {
		s = append(s, cur)
		if cur.Next != nil {
			break
		}
		cur = cur.Next
	}
	delNode := s[len(s)-n]
	if delNode.Next != nil {
		delNode.Val = delNode.Next.Val
		delNode.Next = delNode.Next.Next
	} else {
		if len(s) >= 2 {
			preNode := s[len(s)-n-1]
			preNode.Next = nil
		} else {
			return nil
		}
	}
	return head
}

func main() {

}
