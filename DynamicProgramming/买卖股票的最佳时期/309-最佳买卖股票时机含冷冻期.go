package main

/*309-最佳买卖股票时机含冷冻期*/

func maxProfit309(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	dp := make([][]int, n)
	status := make([]int, n*4)
	for i := range dp {
		dp[i] = status[:4]
		status = status[4:]
	}
	dp[0][0] = -prices[0]
	/*
		0：持有股票，今天买入
		1：保持卖出股票状态
		2：今天卖出股票
		3：冷冻期
	*/
	for i := 1; i < n; i++ {
		dp[i][0] = max309(dp[i-1][0], max309(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max309(dp[i-1][1], dp[i-1][3]) //
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max309(dp[n-1][1], max309(dp[n-1][2], dp[n-1][3]))
}

func max309(a, b int) int {
	if a > b {
		return a
	}
	return b
}
