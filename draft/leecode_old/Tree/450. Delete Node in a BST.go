package main

import . "leecode/util"

/*Given a root node reference of a BST and a key, delete the node with the given key in the BST.
Return the root node reference (possibly updated) of the BST.
Basically, the deletion can be divided into two stages:
Search for a node to remove.
If the node is found, delete the node.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 返回值为删除节点后的传入树的根节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return deleteRoot(root)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root

}

// 删除的节点为根节点，返回新的根节点
func deleteRoot(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}
	temp := node.Left
	for nil != temp.Right {
		temp = temp.Right
	}
	temp.Left = deleteNode(node, temp.Val).Left // 从待删除的节点删除前驱，
	temp.Right = node.Right
	return temp
}
