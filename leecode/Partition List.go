package main

import "fmt"
import . "leecode/util"

/*
Given the head of a linked list and a value x, partition it such that all nodes less
than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.
*/

func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{Next: nil}
	greater := &ListNode{Next: nil}
	preLess := less
	preGreater := greater
	for nil != head {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}
	less.Next = preGreater.Next
	greater.Next = nil
	return preLess.Next
}

func main() {
	ListNode6 := &ListNode{Val: 2, Next: nil}
	ListNode5 := &ListNode{Val: 5, Next: ListNode6}
	ListNode4 := &ListNode{Val: 2, Next: ListNode5}
	ListNode3 := &ListNode{Val: 3, Next: ListNode4}
	ListNode2 := &ListNode{Val: 4, Next: ListNode3}
	ListNode1 := &ListNode{Val: 1, Next: ListNode2}
	PrintfListNode(ListNode1)
	fmt.Println("========")
	PrintfListNode(partition(ListNode1, 3))

}
