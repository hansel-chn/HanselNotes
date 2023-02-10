package main

import . "leecode/util"

//给你一个二叉树的根节点root，检查它是否轴对称。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var recursion func(p, q *TreeNode) bool
	recursion = func(p, q *TreeNode) bool {
		if nil == p && nil == q {
			return true
		}
		if nil == p || nil == q {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return recursion(p.Left, q.Right) && recursion(p.Right, q.Left)
	}
	return recursion(root, root)
}

func main() {

}
