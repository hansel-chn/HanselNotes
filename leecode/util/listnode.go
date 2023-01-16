package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintfListNode(node *ListNode) {
	for nil != node {
		fmt.Println(node.Val)
		node = node.Next
	}
}
