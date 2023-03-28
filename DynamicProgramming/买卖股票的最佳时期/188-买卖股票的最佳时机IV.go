package main

/*188-买卖股票的最佳时机IV*/
// 买卖股票的最佳时机IV 动态规划
// 时间复杂度O(kn) 空间复杂度O(kn)
func maxProfit188(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	status := make([]int, (2*k+1)*len(prices))
	for i := range dp {
		dp[i] = status[:2*k+1]
		status = status[2*k+1:]
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max188(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max188(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}
