package main

/*33-搜索旋转排序数组*/

func search(nums []int, target int) int {
	l := len(nums) - 1

	if target < nums[0] && target > nums[l] {
		return -1
	}

	if target >= nums[0] {
		for i := 0; i <= l; i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := l; i >= 0; i++ {
			if nums[i] == target {
				return i
			}
		}
	}

	return -1
}
