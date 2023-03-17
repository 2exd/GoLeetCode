package main

/*164-最大间距*/
func findMax164(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumGap(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	// m is the maximal number in nums
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		m = findMax164(m, nums[i])
	}
	exp := 1 // 1, 10, 100, 1000 ...
	R := 10  // 10 digits
	// use to store sorted nums
	aux := make([]int, len(nums))
	for (m / exp) > 0 { // Go through all digits from LSB to MSB
		// 保存数组中每个数字最后末尾值
		count := make([]int, R)
		for i := 0; i < len(nums); i++ {
			count[(nums[i]/exp)%10]++
		}
		// 更改count[i]。目的是让更改后的count[i]的值，是该数据在aux[]中的位置。
		for i := 1; i < len(count); i++ {
			count[i] += count[i-1]
		}
		// 将数据存储到辅助数组aux[]中
		for i := len(nums) - 1; i >= 0; i-- {
			tmp := count[(nums[i]/exp)%10]
			tmp--
			aux[tmp] = nums[i]
			count[(nums[i]/exp)%10] = tmp
		}
		// 将排序好的数据赋值给nums[]
		for i := 0; i < len(nums); i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}
	maxValue := 0
	for i := 1; i < len(aux); i++ {
		maxValue = findMax164(maxValue, aux[i]-aux[i-1])
	}
	return maxValue
}

// func main() {
// 	maximumGap([]int{42, 78, 9, 23, 56})
// }
