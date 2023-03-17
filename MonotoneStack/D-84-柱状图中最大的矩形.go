package main

func largestRectangleArea(heights []int) int {
	// 声明max并初始化为0
	max := 0
	// 使用切片实现栈
	stack := make([]int, 0)
	// 数组头部加入0
	heights = append([]int{0}, heights...)
	// 数组尾部加入0
	heights = append(heights, 0)
	// 初始化栈，序号从0开始
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		// 结束循环条件为：当即将入栈元素>top元素，也就是形成非单调递增的趋势
		for heights[stack[len(stack)-1]] > heights[i] {
			// mid 是top
			mid := stack[len(stack)-1]
			// 出栈
			stack = stack[0 : len(stack)-1]
			// left是top的下一位元素，i是将要入栈的元素
			left := stack[len(stack)-1]
			// 高度x宽度
			tmp := heights[mid] * (i - left - 1)
			if tmp > max {
				max = tmp
			}
		}
		stack = append(stack, i)
	}
	return max
}
