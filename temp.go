package main

func lastStoneWeightII(stones []int) int {
	dp := make([]int, 15001)

	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2

	for i := 1; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
