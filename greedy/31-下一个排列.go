package greedy

/*31-下一个排列*/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l := len(nums)
	i, j, k := l-2, l-1, l-1

	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for x, y := j, l-1; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}

}
