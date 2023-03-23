package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	mapNum := map[int]int{}
	// 记录每个元素出现的次数
	for _, item := range nums {
		mapNum[item]++
	}

	h := &IntHeap347{}
	heap.Init(h)
	// 所有元素入堆，堆的长度为k
	for key, value := range mapNum {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	// 按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

type IntHeap347 [][2]int

func (h IntHeap347) Len() int { return len(h) }

func (h IntHeap347) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h IntHeap347) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap347) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap347) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
