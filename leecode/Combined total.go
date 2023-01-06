package main

import (
	"fmt"
	"sort"
)

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

/*
俩条路经一条target值变，另一条index值变，每个节点有相对应的target, index, combination状态
combination状态放在外面，其实就是节点没存值，另外两个也可以，见combinationSum4
区别就在于每次存值还是每次运算
*/
func combinationSum1(candidates []int, target int) [][]int {
	combination := make([]int, 0)
	result := make([][]int, 0)
	var des func(target int, idx int)
	des = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}
		if 0 > target {
			return
		}
		combination = append(combination, candidates[idx])
		des(target-candidates[idx], idx)
		combination = combination[:len(combination)-1]
		des(target, idx+1)
		return
	}
	des(target, 0)
	return result
}

/*
和combinationSum1相比，代码区别只是顺序区别；过程思考不同。

以[1,2,3,4,5]为例，为防止结果重复，index递增。combinationSum1遍历思路从树上看是：
从根节点，左子树遍历完，获得的slice结果必定contain 1(组合结果为任意个数1(个数大于等于1)，与2，3，4，5，6的排列组合)，
右子树剩余结果必定not contain 1(index递增)。依次切分左子树易于理解。

combinationSum2不容易理解
*/
func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combination := make([]int, 0)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if idx == len(candidates) {
			return
		}
		if 0 > target {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}
		dfs(idx+1, target)
		combination = append(combination, candidates[idx])
		dfs(idx, target-candidates[idx])
		combination = combination[:len(combination)-1]
		return
	}
	dfs(0, target)
	return result
}

func combinationSum3(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combination := make([]int, 0)
	sort.Ints(candidates)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if idx == len(candidates) {
			return
		}
		if 0 > target {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
		}
		combination = append(combination, candidates[idx])
		dfs(idx, target-candidates[idx])
		combination = combination[:len(combination)-1]

		if target < candidates[idx] {
			return
		}
		dfs(idx+1, target)
		return
	}
	dfs(0, target)
	return result
}

func combinationSum4(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combination := make([]int, 0)
	idx := 0
	sort.Ints(candidates)
	var dfs func()
	dfs = func() {
		if idx == len(candidates) {
			return
		}
		if 0 > target {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
		}
		combination = append(combination, candidates[idx])
		target = target - candidates[idx]
		dfs()
		target = target + candidates[idx]
		combination = combination[:len(combination)-1]

		if target < candidates[idx] {
			return
		}

		idx++
		dfs()
		idx--
		return
	}
	dfs()
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	results := combinationSum4(candidates, target)
	fmt.Println(results)
}
