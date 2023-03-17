package main

import (
	"container/heap"
)

// func main() {
// 	var nums = []int{6, 1, 5, 3, 8, 7, 2, 4}
// 	// HeapSort2(nums)
// 	h := IntHeap(nums)
// 	heap.Push(&h, 9)
// 	fmt.Println(h)
// }
func HeapSort2(values []int) {
	h := IntHeap(values)
	heap.Init(&h)
	n := h.Len()
	for i := 0; i < n; i++ {
		heap.Pop(&h)
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// 大顶堆
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop `Pop` 之所以这样实现，是因为在heap包的源码中，`Pop` 在调用 `IntHeap` 的 `Pop` 之前进行了一些操作
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
