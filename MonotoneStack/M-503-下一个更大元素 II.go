package main

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	result := make([]int, length, length)
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}
	// 单调递减，存储数组下标索引
	stack := make([]int, 0)
	for i := 0; i < length*2; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			result[index] = nums[i%length]
		}
		stack = append(stack, i%length)
	}
	return result
}
