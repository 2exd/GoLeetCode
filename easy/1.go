package main

/*1.两数之和*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		// if nums[i] > target {
		//	continue
		// }
		temp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if temp == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*复杂度分析
时间复杂度：O(N^2)，其中 N 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)。
。*/

/*第一名的做法, 哈希表, 耗时少*/
func twoSumNo1(nums []int, target int) []int {
	// 创建map
	m := make(map[int32]int16)
	for i := range nums {
		// 如果map中已有target-nums[i]的值就返回
		if index, ok := m[int32(target-nums[i])]; ok {
			return []int{int(index), i}
		}
		m[int32(nums[i])] = int16(i)
	}
	return nil
}

/*复杂度分析
时间复杂度：O(N)，其中 N 是数组中的元素数量。对于每一个元素 x，我们可以 O(1) 地寻找 target - x。
空间复杂度：O(N)，其中 N 是数组中的元素数量。主要为哈希表的开销。。*/

// func main() {
//	var arr = []int{-1,-2,-3,-4,-5}
//	fmt.Println(twoSumNo1(arr,-8))
// }
