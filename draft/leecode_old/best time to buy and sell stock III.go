package main

import "fmt"

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

Find the maximum profit you can achieve. You may complete at most two transactions.

Note: You may not engage in multiple transactions simultaneously
(i.e., you must sell the stock before you buy again).
*/

/*
思路1：
可以利用maxProfit找出一次最大的收益，得到最大最小position，使用for循环对整个切片进行遍历，
1. 当idx小于最小positon，寻找0-idx的最大收益即可
2. 当idx大于最大position，寻找idx-len（prices）的最大收益即可
3. idx在最小position和最小position之间，遍历
* 可以存在只买一次，用0来判断区间收益不能为负
思路2：dp
*/
func maxProfit3(prices []int) int {
	buy1 := -prices[0]
	buy2 := -prices[0]
	sell1 := 0
	sell2 := 0
	for i := 1; i < len(prices); i++ {
		buy1 = max3(buy1, -prices[i])
		sell1 = max3(sell1, buy1+prices[i])
		buy2 = max3(buy2, sell1-prices[i])
		sell2 = max3(sell2, buy2+prices[i])
	}
	return sell2
}

func max3(value1 int, value2 int) int {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}

func main() {
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit3(prices1))
	fmt.Println(maxProfit3(prices2))
	fmt.Println(maxProfit3(prices3))
}
