package main

/*55. 跳跃游戏*/

// 贪心算法
func canJUmp(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	coverRange := 0
	for i := 0; i < coverRange; i++ {
		coverRange = max(coverRange, nums[i]+i)
		if coverRange >= len(nums)-1 {
			return true
		}
	}
	return false
}
