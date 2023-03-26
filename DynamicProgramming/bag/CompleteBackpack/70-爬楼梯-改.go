package main

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	m := 2
	// 先遍历背包 => 楼梯
	for i := 1; i <= n; i++ {
		// 在遍历物品 => 步长
		for j := 0; j < m; j++ {
			if j > i {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}
