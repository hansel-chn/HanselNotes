package main

import "fmt"

/*
Given an unsorted integer array nums, return the smallest missing positive integer.
You must implement an algorithm that runs in O(n) time and uses constant extra space.
*/

func firstMissingPositive(nums []int) int {
	numsLength := len(nums)
	for i, _ := range nums {
		for nums[i] <= numsLength && nums[i] > 0 && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, num := range nums {
		if i != num-1 {
			return i + 1
		}
	}
	return numsLength + 1
}

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}
