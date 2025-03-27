package main

import "fmt"

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock.
You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
*/

func maxProfit2(prices []int) int {
	idx := 0
	Profit := 0
	for idx = 0; idx < len(prices)-1; idx++ {
		if prices[idx] < prices[idx+1] {
			Profit += prices[idx+1] - prices[idx]
		}
	}
	return Profit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit2(prices1))
	fmt.Println(maxProfit2(prices2))
	fmt.Println(maxProfit2(prices3))
}
