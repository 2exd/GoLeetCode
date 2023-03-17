package main

import "math/rand"

/*398. 随机数索引*/

type Solution map[int][]int

func Constructor398(nums []int) Solution {
	pos := map[int][]int{}
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	return pos
}

func (pos Solution) Pick(target int) int {
	indices := pos[target]
	return indices[rand.Intn(len(indices))]
}

/*
复杂度分析
时间复杂度：初始化为 O(n)，pick 为 O(1)，其中 n 是 nums 的长度。
空间复杂度：O(n)。我们需要 O(n) 的空间存储 n 个下标。
*/
