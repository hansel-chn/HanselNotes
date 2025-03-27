package main

import (
	"fmt"
	. "leecode/util"
)

/*
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a
height-balanced binary search tree.
*/

// func sortedListToBST(head *ListNode) *TreeNode {
//
// }
var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	count := getLength(head)
	return buildTree(1, count)
}

func getLength(head *ListNode) int {
	count := 0
	for nil != head {
		count++
		head = head.Next
	}
	return count
}

func buildTree(left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	treeNode := &TreeNode{}
	mid := (left + right + 1) / 2
	treeNode.Left = buildTree(left, mid-1)
	treeNode.Val = globalHead.Val
	globalHead = globalHead.Next
	treeNode.Right = buildTree(mid+1, right)
	return treeNode
}

func main() {
	ListNode5 := &ListNode{Val: 9, Next: nil}
	ListNode4 := &ListNode{Val: 5, Next: ListNode5}
	ListNode3 := &ListNode{Val: 0, Next: ListNode4}
	ListNode2 := &ListNode{Val: -3, Next: ListNode3}
	ListNode1 := &ListNode{Val: -10, Next: ListNode2}
	PrintfListNode(ListNode1)
	fmt.Println("========")
	sortedListToBST(ListNode1)

}
