package main

/*213-打家劫舍II*/
// 时间复杂度O(n) 空间复杂度O(n)
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max213(nums[0], nums[1])
	}

	result1 := robRange(nums, 0)
	result2 := robRange(nums, 1)
	return max213(result1, result2)
}

// 偷盗指定的范围
func robRange(nums []int, start int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[start]

	for i := 2; i < len(nums); i++ {
		dp[i] = max213(dp[i-2]+nums[i-1+start], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}
