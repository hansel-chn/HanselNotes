package main

import (
	"fmt"
	. "leecode/util"
)

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	preHead := &ListNode{Next: head}
	cur := preHead
	for nil != cur.Next && nil != cur.Next.Next {
		if cur.Next.Val == cur.Next.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return preHead.Next
}

func main() {
	ListNode5 := &ListNode{Val: 3, Next: nil}
	ListNode4 := &ListNode{Val: 3, Next: ListNode5}
	ListNode3 := &ListNode{Val: 2, Next: ListNode4}
	ListNode2 := &ListNode{Val: 1, Next: ListNode3}
	ListNode1 := &ListNode{Val: 1, Next: ListNode2}
	PrintfListNode(ListNode1)
	fmt.Println("========")
	PrintfListNode(deleteDuplicates(ListNode1))
}
