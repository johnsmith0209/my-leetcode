/**
敲黑板！
快慢指针找中点
快指针前进速度是慢指针*2 所以当快指针到头的时候 慢指针的位置正好是一半

*/

package main

import "fmt"

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

func isPalindrome(head *ListNode) bool { // 28ms
	cur, arr := head, []*ListNode{}
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	l := len(arr)
	i, j := 0, l-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	i, cur = 0, head
	for i < l {
		if cur.Val != arr[i].Val {
			return false
		}
		i++
		cur = cur.Next
	}
	return true
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

func isPalindrome2(head *ListNode) bool { // 32ms
	if head == nil { // linked list length == 0
		return true
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}
	if slow == nil { // linked list length == 1
		return true
	}
	newHead := &ListNode{Next: nil, Val: slow.Val}
	for slow.Next != nil {
		slow = slow.Next
		newHead = &ListNode{Next: newHead, Val: slow.Val}
	}
	for head != nil && newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head, newHead = head.Next, newHead.Next
	}
	return true
}

func isPalindrome3(head *ListNode) bool { // 32ms 20ms
	if head == nil { // linked list length == 1
		return true
	}
	slow, fast, prev := head, head, &ListNode{}
	prev = nil
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}
	if slow == nil { // linked list length == 1
		return true
	}
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}
	return true
}

func main() {

}
