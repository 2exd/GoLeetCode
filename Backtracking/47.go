package main

import "sort"

var (
	path47 []int
	res47  [][]int

	st47 []bool // state的缩写
)

func dfs47(cur int, nums []int) {
	if cur == len(nums) {
		tmp := make([]int, len(path47))
		copy(tmp, path47)
		res47 = append(res47, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {

		if i != 0 && nums[i] == nums[i-1] && !st47[i-1] { // 去重，用st来判别是深度还是广度
			continue
		}

		if !st47[i] {
			path47 = append(path47, nums[i])
			st47[i] = true
			dfs47(cur+1, nums)
			st47[i] = false
			path47 = path47[:len(path47)-1]
		}

	}
}

func permuteUnique(nums []int) [][]int {
	res47 = make([][]int, 0, len(nums))
	path47 = make([]int, 0, len(nums))
	st47 = make([]bool, len(nums))
	sort.Ints(nums)
	dfs47(0, nums)
	return res47
}
