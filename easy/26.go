package main

func removeDuplicates(nums []int) int {

	length := len(nums)
	res := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}
