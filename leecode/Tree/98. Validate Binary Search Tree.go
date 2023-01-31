package main

import (
	. "leecode/util"
	"math"
)

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	isValid := true
	refer := math.MinInt64
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if !isValid {
			return
		}
		if refer >= node.Val {
			isValid = false
			return
		} else {
			refer = node.Val
		}
		traversal(node.Right)
	}
	traversal(root)
	return isValid
}

func main() {
	//isValidBST()
}
