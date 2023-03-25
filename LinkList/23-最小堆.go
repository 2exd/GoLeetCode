package main

import "container/heap"

// func main() {
//
// }

func mergeHeapKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	p := dummy
	var h nodeHeap
	heap.Init(&h)
	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}
	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next
}

type nodeHeap []*ListNode

func (heap nodeHeap) Len() int { return len(heap) }

func (heap nodeHeap) Less(i, j int) bool {
	return heap[i].Val < heap[j].Val
}

func (heap *nodeHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *nodeHeap) Push(x interface{}) {
	*heap = append(*heap, x.(*ListNode))
}

func (heap *nodeHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[:n-1]
	return x
}
