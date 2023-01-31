package main

import "fmt"

/*
Given an integer n, return the number of structurally unique BST's (binary search trees)
which has exactly n nodes of unique values from 1 to n.
*/

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func numTrees1(n int) int {
	var recursion func(n int) (num int)
	recursion = func(n int) (num int) {
		if 0 == n {
			return 1
		}
		if 1 == n {
			return 1
		}
		for i := 1; i < n+1; i++ {
			num += recursion(i-1) * recursion(n-i)
		}
		return num
	}
	return recursion(n)
}

func main() {
	fmt.Println(numTrees1(18))
	fmt.Println(numTrees(18))
}
