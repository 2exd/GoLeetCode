package main

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//func main() {
//	var arr = []int{0, 1, 0, 3, 12}
//	moveZeroes(arr)
//}
