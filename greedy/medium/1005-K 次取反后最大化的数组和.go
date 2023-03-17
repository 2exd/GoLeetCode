package main

import (
	"math"
	"sort"
)

/*1005. K 次取反后最大化的数组和*/
func myLargestSumAfterKNegations(nums []int, k int) int {
	absNums := QuickSortAbs(nums)
	for i := 0; i < len(absNums); i++ {
		if absNums[i] < 0 && k > 0 {
			absNums[i] *= -1
			k--
		}
	}
	if k%2 == 1 {
		absNums[len(absNums)-1] *= -1
	}
	sum := 0
	for _, num := range absNums {
		sum += num
	}
	return sum
}

func QuickSortAbs(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]          // 第一个数据
	low := make([]int, 0, 0)     // 比我小的数据
	high := make([]int, 0, 0)    // 比我大的数据
	mid := make([]int, 0, 0)     // 与我一样大的数据
	mid = append(mid, splitData) // 加入一个
	for i := 1; i < len(arr); i++ {
		if myAbs(arr[i]) > myAbs(splitData) {
			low = append(low, arr[i])
		} else if myAbs(arr[i]) < myAbs(splitData) {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSortAbs(low), QuickSortAbs(high)
	myArr := append(append(low, mid...), high...)
	return myArr
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func largestSumAfterKNegations(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}
	// 可以多次变换同一个下标，因此选择最小的数反复变换
	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// func main() {
// 	nums := []int{3, -1, 0, 2}
// 	negations := largestSumAfterKNegations(nums, 3)
// 	fmt.Println(negations)
// }
