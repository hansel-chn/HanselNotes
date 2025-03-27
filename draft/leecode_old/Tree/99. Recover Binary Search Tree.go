package main

import . "leecode/util"

/*You are given the root of a binary search tree (BST),
where the values of exactly two nodes of the tree were swapped by mistake.
Recover the tree without changing its structure.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	var x, y, pre *TreeNode
	var recursion func(root *TreeNode)
	flag := 0

	// 如果pre不太好取-inf，可以声明个结构体，此时地址是nil，可以通过判断nil合并到一起（比如判断大小但是初始值没办法用inf定义）
	recursion = func(root *TreeNode) {
		if nil == root {
			return
		}
		recursion(root.Left)
		if flag == 1 {
			return
		}
		if pre != nil && pre.Val > root.Val {
			y = root
			if nil == x {
				x = pre
			} else {
				flag = 1
				return
			}
		}
		pre = root
		recursion(root.Right)
	}
	recursion(root)
	x.Val, y.Val = y.Val, x.Val
}

func main() {
	//num := -123456
	//str := strconv.Itoa(num)
	//for _, i2 := range str {
	//	//fmt.Println(i2)
	//}
	//a := math.MaxFloat64
	//b := math.MaxInt64
	//fmt.Println(math.Inf(1))
	//fmt.Println(math.Inf(1) > b)
	//
	//var c int32 = 35
	//fmt.Println(c > b)
	//math.min
	//fmt.Println(strconv.IntSize)

	//var num int
	//fmt.Println(num)
}
