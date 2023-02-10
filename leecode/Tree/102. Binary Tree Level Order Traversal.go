package main

import . "leecode/util"

/*给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	rlt := make([][]int, 0)
	var traversal func(node *TreeNode, idx int)
	traversal = func(node *TreeNode, idx int) {
		if nil == node {
			return
		}
		if len(rlt) <= idx {
			rlt = append(rlt, []int{})
		}
		rlt[idx] = append(rlt[idx], node.Val)
		traversal(node.Left, idx+1)
		traversal(node.Right, idx+1)
	}
	traversal(root, 0)
	return rlt
}
