package main

/*
377-组合总和 Ⅳ
*/
func combinationSum4(nums []int, target int) int {
	// 定义dp数组
	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1
	// 注意，此题顺序不同的序列被视作不同的组合。
	// 因此先遍历背包,再循环遍历物品。
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

//
// func main() {
// 	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
// }
