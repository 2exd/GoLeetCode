package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	// `+` 的元素和为 x，`-` 的元素和为 sum-x =>  x - (sum - x) = target => x = (target + sum) / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs494(target) > sum {
		return 0
	}
	// x 不存在
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bagSize := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for _, num := range nums {
		for j := bagSize; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[bagSize]
}
func abs494(x int) int {
	return int(math.Abs(float64(x)))
}
