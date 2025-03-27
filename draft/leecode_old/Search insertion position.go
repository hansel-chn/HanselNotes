package main

import "fmt"

func searchInsert(nums []int, target int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex <= rightIndex {
		tempIndex := (rightIndex + leftIndex) / 2
		if nums[tempIndex] == target {
			return tempIndex
		} else if nums[tempIndex] > target {
			rightIndex = tempIndex - 1
		} else {
			leftIndex = tempIndex + 1
		}
	}
	return leftIndex
}

//func searchInsert(nums []int, target int) int {
//	n := len(nums)
//	left, right := 0, n-1
//	ans := n
//	for left <= right {
//		mid := (right-left)>>1 + left
//		if target <= nums[mid] {
//			ans = mid
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return ans
//}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7

	index := searchInsert(nums, target)
	fmt.Println(index)
}
