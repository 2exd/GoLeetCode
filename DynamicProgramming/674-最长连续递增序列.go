package main

/*674-最长连续递增序列*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, count := 1, 1
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i+1] > nums[i] {
			count++
		} else {
			count = 1
		}
		if res < count {
			res = count
		}
	}
	return res
}
