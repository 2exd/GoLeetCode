package main

/*122-买卖股票的最佳时期III*/
func maxProfit123(prices []int) int {
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < length; i++ {
		dp[i][0] = dp[i-1][0]                               // 无状态，这里可优化
		dp[i][1] = max123(dp[i-1][1], dp[i-1][0]-prices[i]) // 第一次买
		dp[i][2] = max123(dp[i-1][2], dp[i-1][1]+prices[i]) // 第一次卖
		dp[i][3] = max123(dp[i-1][3], dp[i-1][2]-prices[i]) // 第二次买
		dp[i][4] = max123(dp[i-1][4], dp[i-1][3]+prices[i]) // 第二次卖
	}
	return dp[length-1][4]
}

func max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}
