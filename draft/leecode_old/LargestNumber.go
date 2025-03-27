package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.
*/

func largestNumber(nums []int) string {
	numStr := make([]string, 0)
	//fmt.Println(numStr)
	for _, num := range nums {
		numStr = append(numStr, strconv.Itoa(num))
	}
	//fmt.Println(numStr)
	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] >= numStr[j]+numStr[i]
	})
	//sort.Sort(sort.Reverse(sort.StringSlice(numStr)))
	//fmt.Println(numStr)

	if numStr[0] == "0" {
		return "0"
	}
	//k := 0
	//for k < len(nums)-1 && numStr[0] == "0" {
	//	numStr = numStr[1:]
	//	//fmt.Println(len(numStr))
	//	//numStr = append(numStr, "0")
	//	k++
	//}
	rlt := strings.Join(numStr, "")
	return rlt
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(largestNumber(nums))
	//fmt.Println(strconv.IntSize)
	//fmt.Println(strconv.Itoa(1))
	//fmt.Println(string(rune(25105)))
	//fmt.Println([]rune("我")[0])
	//numStr := []string{"20", "2", "23"}
	//sort.Strings(numStr)
	//fmt.Println(numStr)
	//sort.Sort(sort.Reverse(sort.StringSlice(numStr)))
	//fmt.Println(numStr)
}
