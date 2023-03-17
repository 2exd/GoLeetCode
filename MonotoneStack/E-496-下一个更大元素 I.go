package main

func myNextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	distance := make(map[int]int, len(nums2))
	var stack []int

	for i, v := range nums2 {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > nums2[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// distance
			distance[nums2[top]] = nums2[i]
		}
		// push
		stack = append(stack, i)
	}

	for i := 0; i < len(stack); i++ {
		distance[nums2[stack[i]]] = -1
	}

	for i := 0; i < len(nums1); i++ {
		ans[i] = distance[nums1[i]]
	}

	return ans
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	mp := map[int]int{}
	for i, v := range nums1 {
		mp[v] = i
	}
	// 单调栈
	stack := []int{}
	stack = append(stack, 0)

	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {

			top := stack[len(stack)-1]

			if _, ok := mp[nums2[top]]; ok { // 看map里是否存在这个元素
				index := mp[nums2[top]] // 根据map找到nums2[top] 在 nums1中的下标
				res[index] = nums2[i]
			}

			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, i)
	}
	return res
}
