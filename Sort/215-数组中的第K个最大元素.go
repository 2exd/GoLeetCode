package main

import (
	"math/rand"
	"time"
)

/*
LFUCache5-数组中的第K个最大元素
*/

// 快速排序

func findKthLargestQuickSelect(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) (res int) {
	q := randPartition(nums, l, r)
	if q == k {
		return nums[q]
	} else if q < k {
		return quickSelect(nums, q+1, r, k)
	} else {
		return quickSelect(nums, l, q-1, k)
	}
}

func randPartition(nums []int, l, r int) (res int) {
	i := rand.Int()%(r-l+1) + 1
	// exchange
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	x := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+i]
	return i + 1
}

// 堆排序
func findKthLargestHeapSelect(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {

	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r := 2*i+1, 2*i+2
	largest := i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
