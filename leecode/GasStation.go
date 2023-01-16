package main

import "fmt"

/*
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank, and it costs 'cost[i]' of gas to travel from the ith station
to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel
around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution,
it is guaranteed to be unique
*/

func canCompleteCircuit1(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		position := i
		residual := gas[i] - cost[i]
		count := 0
		for residual >= 0 {
			count++
			if len(gas) <= count {
				return position
			}
			i++
			residual = residual + gas[i%len(gas)] - cost[i%len(gas)]
		}
		i++
	}
	return -1
}

// 在最后加跳到下一个索引，索引和值不对应，不需要出来以后++。canCompleteCircuit1 循环内索引和值对应，所以出来还得++，且第一个residual在外面生成
func canCompleteCircuit2(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		position := i
		residual := 0
		count := 0
		for residual >= 0 {
			if len(gas) <= count {
				return position
			}
			residual = residual + gas[i%len(gas)] - cost[i%len(gas)]
			i++
			count++
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit2(gas, cost))
}
