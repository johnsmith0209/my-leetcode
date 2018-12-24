package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head, l1Cur, l2Cur := &ListNode{Val: 0, Next: nil}, l1, l2
	newCur, isHead := head, true
	if l1 == nil && l2 == nil {
		return nil
	}
	for l1Cur != nil || l2Cur != nil {
		next, curList := l1Cur, 1
		if l1Cur == nil && l2Cur != nil {
			fmt.Println("list1 empty, append list 2 into new", l2Cur.Val)
			next, curList = l2Cur, 2
		} else if l1Cur != nil && l2Cur == nil {
			fmt.Println("list2 empty, append list 1 into new", l1Cur.Val)
			next, curList = l1Cur, 1
		} else {
			if l1Cur.Val >= l2Cur.Val {
				fmt.Println("list1 >= list2, append list 2 into new", l2Cur.Val)
				next, curList = l2Cur, 2
			} else {
				fmt.Println("list2 > list, append list 1 into new", l1Cur.Val)
				next, curList = l1Cur, 1
			}
		}
		if isHead {
			newCur.Val = next.Val
			isHead = false
		} else {
			temp := &ListNode{Val: next.Val, Next: nil}
			newCur.Next = temp
			newCur = newCur.Next
		}
		if curList == 1 {
			fmt.Println("list1 got next")
			l1Cur = l1Cur.Next
		} else {
			fmt.Println("list2 got next")
			l2Cur = l2Cur.Next
		}
	}
	return head
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
	// a := ListNode{Val: 1, Next: nil}
	d, e, f := ListNode{Val: 1, Next: nil}, ListNode{Val: 3, Next: nil}, ListNode{Val: 6, Next: nil}
	a.Next = &b
	b.Next = &c
	d.Next = &e
	e.Next = &f
	rzt := mergeTwoLists(&a, &d)
	printList(rzt)
}
