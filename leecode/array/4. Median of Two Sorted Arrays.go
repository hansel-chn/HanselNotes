package main

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if 0 == (m+n-1)%2 {
		return searchIdx(nums1, nums2, (m+n-1)/2)
	} else {
		return (searchIdx(nums1, nums2, (m+n-1)/2) + searchIdx(nums1, nums2, (m+n-1)/2+1)) / 2.0
	}
}

func searchIdx(nums1 []int, nums2 []int, k int) (val float64) {
	num1Idx := 0
	num2Idx := 0
	for i := 0; i <= k; i++ {
		if num1Idx == len(nums1) {
			return float64(nums2[k-i+num2Idx])
		}

		if num2Idx == len(nums2) {
			return float64(nums1[k-i+num1Idx])
		}

		if nums1[num1Idx] <= nums2[num2Idx] {
			val = float64(nums1[num1Idx])
			num1Idx++
		} else {
			val = float64(nums2[num2Idx])
			num2Idx++
		}
	}
	return val
}

func separateFindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if 0 == (m+n+1)%2 {
		return searchIdx(nums1, nums2, (m+n+1)/2)
	} else {
		return (searchIdx(nums1, nums2, (m+n+1)/2) + searchIdx(nums1, nums2, (m+n+1)/2+1)) / 2.0
	}
}
func separateSearchIdx(nums1 []int, nums2 []int, k int) (val float64) {
	num1Idx := 0
	num2Idx := 0
	for k > 0 {
		if num1Idx == len(nums1) {
			return float64(nums2[num2Idx+k-1])
		}
		if num2Idx == len(nums2) {
			return float64(nums1[num1Idx+k-1])
		}

		if k == 1 {
			return float64(min(nums1[num1Idx], nums2[num2Idx]))
		}
		half := min(min(k/2-1, len(nums1)-1-num1Idx), len(nums2)-1-num2Idx)

		newIndex1 := num1Idx + half
		newIndex2 := num2Idx + half

		//newIndex1 := min(num1Idx+half, len(nums1)-1)
		//newIndex2 := min(num2Idx+half, len(nums2)-1)

		if nums1[newIndex1] <= nums2[newIndex2] {
			num1Idx = newIndex1 + 1
		} else {
			num2Idx = newIndex2 + 1
		}
		k = k - half - 1
	}

	return val
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
