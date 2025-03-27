package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

// 虽然优化了但还是三重循环，最坏情况每次第三重循环的结果都在最后
func threeSumFault(nums []int) [][]int {
	sort.Ints(nums)
	rlt := make([][]int, 0)
	for i := 0; i < len(nums); {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < len(nums); {
			target := -nums[i] - nums[j]
			if target < nums[j] {
				break
			}

			for k := j + 1; k < len(nums); {
				if target < nums[k] {
					break
				}
				if target == nums[k] {
					rlt = append(rlt, []int{nums[i], nums[j], nums[k]})
					break
				}
				for k+1 != len(nums) && nums[k] == nums[k+1] {
					k = k + 1
				}
				k = k + 1
			}
			for j+1 != len(nums) && nums[j] == nums[j+1] {
				j = j + 1
			}
			j = j + 1
		}
		for i+1 != len(nums) && nums[i] == nums[i+1] {
			i = i + 1
		}
		i = i + 1
	}
	return rlt
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	rlt := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		tempIdx := len(nums) - 1

		for j := i + 1; j < len(nums); j++ {
			if tempIdx <= j {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := tempIdx; nums[k]+nums[j] >= target && j < k; k-- {
				if k < tempIdx && nums[k] == nums[k+1] {
					tempIdx = k
					continue
				}
				if target == nums[k]+nums[j] {
					rlt = append(rlt, []int{nums[i], nums[j], nums[k]})
				}
				tempIdx = k
			}
		}
	}
	return rlt
}

//func threeSumHash(nums []int) [][]int {
//	sort.Ints(nums)
//	rlt := make([][]int, 0)
//	for i := 0; i < len(nums); i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		if nums[i] > 0 {
//			break
//		}
//		target := -nums[i]
//		tempIdx := len(nums) - 1
//		//hash := make(map[int]int)
//
//		for j := i + 1; j < len(nums); j++ {
//			if tempIdx <= j {
//				break
//			}
//			if j > i+1 && nums[j] == nums[j-1] {
//				continue
//			}
//			for k := tempIdx; nums[k]+nums[j] >= target && j < k; k-- {
//				if k < tempIdx && nums[k] == nums[k+1] {
//					tempIdx = k
//					continue
//				}
//				if target == nums[k]+nums[j] {
//					rlt = append(rlt, []int{nums[i], nums[j], nums[k]})
//				}
//				tempIdx = k
//			}
//		}
//	}
//	return rlt
//}

func main() {
	//a := 1
	//b := 2
	//c := []int{a, b}
	//fmt.Println(c)
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
