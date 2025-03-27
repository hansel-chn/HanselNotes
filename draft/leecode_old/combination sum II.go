package main

import (
	"fmt"
	"sort"
)

/*Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.
Each number in candidates may only be used once in the combination.
Note: The solution set must not contain duplicate combinations.*/

/*
基于combinationSum，
*/
func combinationSumIIV1(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	//sortedCandidates := make([]int, 0)
	//sortedCandidates = append(sortedCandidates, candidates[0])
	//for i := 0; i < len(candidates)-1; i++ {
	//	if candidates[i] != candidates[i+1] {
	//		sortedCandidates = append(sortedCandidates, candidates[i+1])
	//	}
	//}
	combination := make([]int, 0)
	rlt := make([][]int, 0)
	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if 0 > target {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			rlt = append(rlt, combinationCopy)
			return
		}
		if idx == len(candidates) {
			return
		}

		combination = append(combination, candidates[idx])
		dfs(target-candidates[idx], idx+1)
		combination = combination[:len(combination)-1]

		if target < candidates[idx] {
			return
		}

		for i := idx; i < len(candidates)-1; i++ { // 核心代码，与combination Sum不同之处
			if candidates[i] == candidates[i+1] {
				continue
			}
			dfs(target, i+1)
			break
		}
	}

	dfs(target, 0)

	return rlt
}

func combinationSumIIV2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	rlt := make([][]int, 0)
	processedCandi := make([][2]int, 0)
	for _, candidate := range candidates {
		if len(processedCandi) == 0 || candidate != processedCandi[len(processedCandi)-1][0] {
			processedCandi = append(processedCandi, [2]int{candidate, 1})
		} else {
			processedCandi[len(processedCandi)-1][1]++
		}
	}

	combination := make([]int, 0)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if 0 > target {
			return
		}
		if 0 == target {
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			rlt = append(rlt, combinationCopy)
			return
		}
		if idx == len(processedCandi) || target < processedCandi[idx][0] {
			return
		}

		var most int
		if target/processedCandi[idx][0] > processedCandi[idx][1] {
			most = processedCandi[idx][1]
		} else {
			most = target / processedCandi[idx][0]
		}
		for i := 1; i <= most; i++ {
			combination = append(combination, processedCandi[idx][0])
			dfs(idx+1, target-i*processedCandi[idx][0])
		}

		combination = combination[:len(combination)-most]

		dfs(idx+1, target)
	}
	dfs(0, target)

	return rlt
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	results := combinationSumIIV2(candidates, target)
	fmt.Println(results)
}
