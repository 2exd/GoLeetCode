package main

func lastStoneWeightII(stones []int) int {
	// 15001 = 30 * 1000 /2 +1
	dp := make([]int, 15001)
	// 求target
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	// 遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			// 推导公式
			dp[j] = max1049(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]
}

func max1049(a, b int) int {
	if a > b {
		return a
	}
	return b
}
