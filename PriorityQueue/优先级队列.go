package main

type ListHeap1 [][]int

func (h ListHeap1) Len() int           { return len(h) }
func (h ListHeap1) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h ListHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ListHeap1) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

type IntHeap1 []int

func (h IntHeap1) Len() int           { return len(h) }
func (h IntHeap1) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}
