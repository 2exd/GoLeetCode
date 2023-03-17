package main

import (
	"container/heap"
)

/*347. 前 K 个高频元素 medium*/

/*
什么是优先级队列呢？

其实就是一个披着队列外衣的堆，因为优先级队列对外接口只是从队头取元素，从队尾添加元素，再无其他取元素的方式，看起来就是一个队列。
*/
//方法一：小顶堆
// 时间复杂度：O(nlogk)
// 空间复杂度：O(n)
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// IHeap 构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//// 方法二:利用O(logn)排序
//func topKFrequent(nums []int, k int) []int {
//	ans:=[]int{}
//	map_num:=map[int]int{}
//	for _,item:=range nums {
//		map_num[item]++
//	}
//	for key,_:=range map_num{
//		ans=append(ans,key)
//	}
//	//核心思想：排序
//	//可以不用包函数，自己实现快排
//	sort.Slice(ans,func (a,b int)bool{
//		return map_num[ans[a]]>map_num[ans[b]]
//	})
//	return ans[:k]
//}
