package main

/*1. 两数之和 easy*/
// 使用map方式解题，降低时间复杂度
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		// 找另外一个元素下标
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			// map添加元素
			m[val] = index
		}
	}
	return []int{}
}
