package main

/*169-多数元素*/
func majorityElement(nums []int) int {
	major, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
		}
		if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	majorityElement([]int{6, 5, 5})
}
