package main

/*
739. 每日温度
*/

// 单调递减栈
func dailyTemperatures(num []int) []int {
	ans := make([]int, len(num))
	var stack []int
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// distance
			ans[top] = i - top
		}
		// push
		stack = append(stack, i)
	}
	return ans
}
