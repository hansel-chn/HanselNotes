package main

import "fmt"

/*
You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints
of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container,
such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

func maxArea(height []int) int {
	leftPointer, rightPointer := 0, len(height)-1
	maxCapacity := 0
	for (rightPointer - leftPointer) > 0 {
		if height[leftPointer] < height[rightPointer] {
			maxCapacity = max(maxCapacity, height[leftPointer]*(rightPointer-leftPointer))
			leftPointer++
		} else {
			maxCapacity = max(maxCapacity, height[rightPointer]*(rightPointer-leftPointer))
			rightPointer--
		}
	}
	return maxCapacity
}

func max(value1 int, value2 int) int {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}

func main() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height2 := []int{1, 1}
	fmt.Println(maxArea(height1))
	fmt.Println(maxArea(height2))
}
