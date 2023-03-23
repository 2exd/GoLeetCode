package main

import "fmt"

// 堆排序

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		// 每次把最大元素放到堆顶
		HeapSortMax(arr, lastLen)
		// 首尾互换
		if i < length {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}
	return arr
}

func HeapSortMax(arr []int, length int) {
	if length <= 1 {
		return
	}

	depth := length/2 - 1
	// 从尾部开始遍历
	for i := depth; i >= 0; i-- {
		top := i
		l := 2*i + 1
		r := 2*i + 2
		if l < length && arr[l] > arr[top] {
			top = l
		}
		if r < length && arr[r] > arr[top] {
			top = r
		}
		// 根和 max(左，右) 互换位置
		if top != i {
			arr[i], arr[top] = arr[top], arr[i]
		}
	}
}
