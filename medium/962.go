package main

func findMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*962. 最大宽度坡*/

func myMaxWidthRamp(nums []int) int {
	maxLen := 0
	for i := 0; i < len(nums)-1; i++ {
		swLen := 0
		for j := len(nums) - 1; j >= i+1; j-- {
			if nums[j] >= nums[i] {
				swLen = j - i
				break
			}
		}
		maxLen = findMax(maxLen, swLen)
	}
	return maxLen
}

/*
复杂度分析
时间复杂度： O(N^2)，其中 N 是 nums 的长度。超出时间限制。
空间复杂度： O(1)。
*/

/*单调栈*/
func maxWidthRamp(nums []int) (maxLen int) {
	n := len(nums)
	if n == 0 {
		return
	}
	stack := make([]int, n)
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			pos := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxLen = findMax(maxLen, i-pos)
			if len(stack) == 0 {
				break
			}
		}
	}
	return
}

//func main() {
//	//nums := []int{6, 0, 8, 2, 1, 5}
//	nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
//	fmt.Println(maxWidthRamp(nums))
//}
