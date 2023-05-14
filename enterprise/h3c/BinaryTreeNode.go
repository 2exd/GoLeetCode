package main

/*
已知一棵节点个数为 n 的二叉树的中序遍历单调递增，求该二叉树能有多少种树形
*/

func numOfTree(n int) int {
	// 计算左边与右边各有多少种可能
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 左边加右边 当前j为根节点
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
