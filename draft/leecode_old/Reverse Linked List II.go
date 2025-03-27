package main

import "fmt"
import . "leecode/util"

/*
Given the head of a singly linked list and two integers left and right where left <= right,
reverse the nodes of the list from position left to position right, and return the reversed list.
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preHead := &ListNode{Next: head}
	cur := preHead
	for i := 1; i < left; i++ {
		cur = cur.Next
	}
	leftPosition := cur
	cur = cur.Next
	for i := left; i < right; i++ {
		temp := cur.Next
		cur.Next = cur.Next.Next
		temp.Next = leftPosition.Next
		leftPosition.Next = temp
	}

	return preHead.Next
}

func main() {
	ListNode5 := &ListNode{Val: 5, Next: nil}
	ListNode4 := &ListNode{Val: 4, Next: ListNode5}
	ListNode3 := &ListNode{Val: 3, Next: ListNode4}
	ListNode2 := &ListNode{Val: 2, Next: ListNode3}
	ListNode1 := &ListNode{Val: 1, Next: ListNode2}
	PrintfListNode(ListNode1)
	fmt.Println("========")
	PrintfListNode(reverseBetween(ListNode1, 2, 4))
}
