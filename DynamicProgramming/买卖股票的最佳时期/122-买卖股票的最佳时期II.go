package main

/*122-买卖股票的最佳时期II*/
func maxProfit122(prices []int) int {
	length := len(prices)
	dp := make([][]int, length)
	status := make([]int, length*2)
	for i := range dp {
		dp[i] = status[:2]
		status = status[2:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max122(dp[i-1][0], dp[i-1][1]-prices[i]) // 买了卖
		dp[i][1] = max122(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max122(a, b int) int {
	if a > b {
		return a
	}
	return b
}
