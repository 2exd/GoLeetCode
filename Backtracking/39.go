package main

import "sort"

var (
	res39  [][]int
	path39 []int
)

func combinationSum(candidates []int, target int) [][]int {
	res39, path39 = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs39(candidates, 0, target)
	return res39
}

func dfs39(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path39))
		copy(tmp, path39)
		res39 = append(res39, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		path39 = append(path39, candidates[i])
		dfs39(candidates, i, target-candidates[i])
		path39 = path39[:len(path39)-1]
	}
}
