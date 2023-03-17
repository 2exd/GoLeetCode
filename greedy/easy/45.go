package main

/*45. 跳跃游戏 II*/

// 贪心算法
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0 // 初始第一格跳跃数一定为0

	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i { // nums[j]为起点,j为往右跳的覆盖范围,这行表示从j能跳到i
				dp[i] = min(dp[j]+1, dp[i]) // 更新最小能到i的跳跃次数
			}
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
