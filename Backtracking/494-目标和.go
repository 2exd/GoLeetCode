package main

var (
	res494  int
	path494 []int
)

func findTargetSumWays(nums []int, target int) int {
	res494 = 0
	path494 = make([]int, 0, len(nums))
	dfs494(nums, target, 0, 0)
	return res494
}

func dfs494(nums []int, target int, start int, sum int) {
	if len(path494) == len(nums) && sum == target {
		res494++
		return
	}

	for i := start; i < len(nums); i++ {
		path494 = append(path494, nums[i])
		dfs494(nums, target, i+1, sum+nums[i])
		dfs494(nums, target, i+1, sum-nums[i])
		path494 = path494[:len(path494)-1]
	}
}

// func main() {
// 	findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
// }
