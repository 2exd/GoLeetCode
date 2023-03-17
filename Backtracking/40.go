package main

import "sort"

var (
	res40  [][]int
	path40 []int
)

func combinationSum2(candidates []int, target int) [][]int {
	res40, path40 = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs40(candidates, 0, target)
	return res40
}

func dfs40(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path40))
		copy(tmp, path40)
		res40 = append(res40, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}

		if i != start && candidates[i-1] == candidates[i] {
			continue
		}

		path40 = append(path40, candidates[i])
		dfs40(candidates, i+1, target-candidates[i])
		path40 = path40[:len(path40)-1]
	}
}
