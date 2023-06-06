package main

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	stack := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		// 大于栈顶，添加
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else if nums[i] < stack[len(stack)-1] {
			// 小于等于栈顶，二分查找
			left, right := 0, len(stack)-1
			for left < right {
				mid := (left + right) / 2
				if nums[i] <= stack[mid] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			stack[left] = nums[i]
		}
	}
	return len(stack)
}
