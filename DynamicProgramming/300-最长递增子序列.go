package main

/*300. 最长递增子序列*/
// 动态规划求解
func lengthOfLIS(nums []int) int {
	// dp数组的定义 dp[i]表示取第i个元素的时候，表示子序列的长度，其中包括 nums[i] 这个元素
	dp := make([]int, len(nums))
	// 初始化，所有的元素都应该初始化为1
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max300(dp[i], dp[j]+1)
			}
		}
		// update ans
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func max300(x, y int) int {
	if x > y {
		return x
	}
	return y
}
