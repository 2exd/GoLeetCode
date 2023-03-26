package main

/*42. 接雨水*/

/*维护一个单调栈，单调栈存储的是下标，满足从栈底到栈顶的下标对应的数组 height 中的元素递减。*/
func trap(height []int) (ans int) {
	// 存下标
	var stack []int
	for i, h := range height {
		// 当前元素>栈尾 弹栈
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			// 栈尾出栈。相当于是桶底坐标
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			// 当前栈尾
			left := stack[len(stack)-1]
			// 弹一次计算一次
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*复杂度分析
时间复杂度：O(n)，其中 n 是数组 height 的长度。从 0 到 n-1 的每个下标最多只会入栈和出栈各一次。
空间复杂度：O(n)，其中 n 是数组 height 的长度。空间复杂度主要取决于栈空间，栈的大小不会超过 n。*/

// func main()  {
//
// 	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
// 	fmt.Println(trap(height))
// }
