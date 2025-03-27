package main

import . "leecode/util"

/*Given an integer n, return all the structurally unique BST's (binary search trees),
which has exactly n nodes of unique values from 1 to n. Return the answer in any order.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var backtrack func(left, right int) []*TreeNode
	backtrack = func(left, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}

		roots := make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			leftRoots := backtrack(left, i-1)
			rightRoots := backtrack(i+1, right)
			for _, leftRoot := range leftRoots {
				for _, rightRoot := range rightRoots {
					root := &TreeNode{Val: i}
					root.Left = leftRoot
					root.Right = rightRoot
					roots = append(roots, root)
				}
			}
		}
		return roots
	}
	return backtrack(1, n)
}

func main() {

}
