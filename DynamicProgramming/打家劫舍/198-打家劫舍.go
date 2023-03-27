package main

/*198-打家劫舍*/
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1) // dp[i]表示偷到第i家能够偷得的最大金额
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max198(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}
