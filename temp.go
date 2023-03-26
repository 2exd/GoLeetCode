package main

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	m := make(map[int]int, 0)
	stack := make([]int, len(nums))
	for i, num := range nums {
		for len(stack) != 0 && num > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m[nums[top]] = num
		}
		stack = append(stack, i)
	}
	for i, num := range nums {
		if m[num] != 0 {
			res[i] = m[num]
		} else {
			res[i] = -1
		}
	}
	return res
}
