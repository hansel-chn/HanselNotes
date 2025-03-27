package main

import "fmt"

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock
and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

func maxProfit(prices []int) int {
	minPosition := 0
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[minPosition] < prices[i] {
			profit = maxValue(profit, prices[i]-prices[minPosition])
		} else {
			minPosition = i
		}
	}
	return profit
}

func maxValue(value1 int, value2 int) int {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	//prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices1))
	//fmt.Println(maxProfit(prices2))
	fmt.Println(maxProfit(prices3))
}
