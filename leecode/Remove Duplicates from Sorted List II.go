package main

import "fmt"

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/

type ListNode struct {
	Val  int
	Next *ListNode
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
}

func printfListNodeRe(node *ListNode) {
	for nil != node {
		fmt.Println(node.Val)
		node = node.Next
	}
}
