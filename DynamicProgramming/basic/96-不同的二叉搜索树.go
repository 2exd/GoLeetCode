package main

func numTrees(n int) int {
	// 1到i为节点组成的二叉搜索树的个数为dp[i]。
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
