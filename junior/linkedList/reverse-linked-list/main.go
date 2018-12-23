package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, next, isHead := head, head.Next, true
	for next != nil {
		if isHead {
			cur.Next = nil
			isHead = false
		}
		fmt.Println("cur is ", cur.Val, "next is ", next.Val)
		// temp := next.Next
		next.Next, cur, next = cur, next, next.Next
		printList(cur)
	}
	return cur
}

func printList(head *ListNode) {
	fmt.Println("-----Print List Start-----")
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Print("Nil\n")
	fmt.Println("-----Print List End-------")
}

func main() {
	a, b, c := ListNode{Val: 1, Next: nil}, ListNode{Val: 2, Next: nil}, ListNode{Val: 3, Next: nil}
	d, e, f := ListNode{Val: 4, Next: nil}, ListNode{Val: 5, Next: nil}, ListNode{Val: 6, Next: nil}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	printList(&a)
	rzt := reverseList(&a)
	printList(rzt)
}
