package main

import . "leecode/util"

//Given the root of a binary tree, return the inorder traversal of its nodes' values.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	rlt := make([]int, 0)
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if nil == node {
			return
		}
		search(node.Left)
		rlt = append(rlt, node.Val)
		search(node.Right)
	}
	search(root)
	return rlt
}

func main() {

}
