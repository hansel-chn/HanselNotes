package main

import "fmt"

/*
You are given a 0-indexed array of integers nums of length n.
You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i.
In other words, if you are at nums[i], you can jump to any nums[i + j] where:

* 0 <= j <= nums[i] and i + j < n
Return the minimum number of jumps to reach nums[n - 1].
The test cases are generated such that you can reach nums[n - 1].
*/

func jump(nums []int) int {
	if 1 == len(nums) {
		return 0
	}
	idx := 0
	count := 1
	maxPosition := nums[idx]
	for maxPosition < len(nums)-1 {
		tmpPosition := 0
		for i := idx + 1; i <= maxPosition; i++ {
			if tmpPosition < i+nums[i] {
				tmpPosition = i + nums[i]
			}
		}
		idx = maxPosition
		maxPosition = tmpPosition
		count++
	}
	return count
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
