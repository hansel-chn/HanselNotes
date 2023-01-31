package main

import (
	"fmt"
	. "leecode/util"
)

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/

func deleteDuplicatesII(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: head}
	cur := preHead
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next.Next
			for tmp.Next != nil && tmp.Val == tmp.Next.Val {
				tmp = tmp.Next
			}
			cur.Next = tmp.Next
		} else {
			cur = cur.Next
		}
	}
	return preHead.Next
}

func main() {
	ListNode7 := &ListNode{Val: 5, Next: nil}
	ListNode6 := &ListNode{Val: 4, Next: ListNode7}
	ListNode5 := &ListNode{Val: 4, Next: ListNode6}
	ListNode4 := &ListNode{Val: 3, Next: ListNode5}
	ListNode3 := &ListNode{Val: 3, Next: ListNode4}
	ListNode2 := &ListNode{Val: 2, Next: ListNode3}
	ListNode1 := &ListNode{Val: 1, Next: ListNode2}

	PrintfListNode(ListNode1)
	fmt.Println("======")
	PrintfListNode(deleteDuplicatesII(ListNode1))
}
