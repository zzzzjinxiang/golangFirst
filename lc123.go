package main

import "math"

func main() {
	n := []int{1,2}
	//dp := make([][2]int,len([2]int{})) 二维数组
	//dp := make([][2]int,len(prices))
	ts := make([][]int, len(n))
	prices1(n)
}

func prices1(prices []int) int {
	dp_0_0 := 0
	dp_0_1 := math.MinInt64
	for i := 0; i < len(prices); i++ {
		dp_0_0 = max(dp_0_0, dp_0_1 + prices[i])
		dp_0_1 = max(dp_0_1, -prices[i])
	}
	return dp_0_0
}

func max(a int, b int) int {
	if (a > b) {
		return a;
	} 
        return b;
}
