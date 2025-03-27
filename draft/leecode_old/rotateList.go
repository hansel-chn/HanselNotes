package main

import (
	"fmt"
	. "leecode/util"
)

/*
Given the head of a linked list, rotate the list to the right by k places.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	//var node *ListNode
	node := &ListNode{}
	node.Next = head

	count := 0
	for nil != node.Next {
		node = node.Next
		count++
	}
	k = k % count

	node = &ListNode{}
	node.Next = head

	for i := 0; i < count-k; i++ {
		node = node.Next
	}

	newHead := &ListNode{}
	newHead.Next = node.Next
	node.Next = nil
	node = newHead
	for i := 0; i < k; i++ {
		node = node.Next
	}
	node.Next = head
	return newHead.Next
}

func main() {
	preHead := new(ListNode)
	node := new(ListNode)

	preHead = node
	//fmt.Println(preHead)
	for i := 1; i < 6; i++ {
		node.Next = new(ListNode)
		node = node.Next
		node.Val = i
	}
	//rotateRight(preHead.Next, 2)
	PrintfListNode(preHead.Next)
	fmt.Println("-=======")
	PrintfListNode(rotateRight(preHead.Next, 2))
}
