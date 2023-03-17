package main

var (
	path491 []int
	res491  [][]int
)

func dfs491(start int, nums []int) {
	if len(path491) >= 2 {
		tmp := make([]int, len(path491))
		copy(tmp, path491)
		res491 = append(res491, tmp)
	}
	used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
	for i := start; i < len(nums); i++ {

		if used[nums[i]] { // 去重
			continue
		}

		if len(path491) == 0 || nums[i] >= path491[len(path491)-1] {
			path491 = append(path491, nums[i])
			used[nums[i]] = true
			dfs491(i+1, nums)
			path491 = path491[:len(path491)-1]
		}

	}
}

func findSubsequences(nums []int) [][]int {
	res491 = make([][]int, 0, len(nums))
	path491 = make([]int, 0, len(nums))
	dfs491(0, nums)
	return res491
}
