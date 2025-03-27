package main

import "fmt"

/*
You are given an integer array nums. You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	idx := 0
	maxPosition := nums[idx]
	for maxPosition < len(nums)-1 {
		temp := 0
		for i := idx + 1; i <= maxPosition; i++ {
			if temp < i+nums[i] {
				temp = i + nums[i]
			}
		}

		if temp <= maxPosition {
			return false
		} else {
			idx = maxPosition
			maxPosition = temp
		}
	}
	return true
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums1))
	fmt.Println(canJump(nums2))
}
