package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Task 结构包含任务的开始和结束时间
type Task struct {
	start, end int
}

// MinHeap 用于实现堆接口的结构
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 和 Pop 用来添加和移除元素
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func taskProcess() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	tasks := make([]Task, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		se := strings.Split(scanner.Text(), " ")
		si, _ := strconv.Atoi(se[0])
		ei, _ := strconv.Atoi(se[1])
		tasks[i] = Task{start: si, end: ei}
	}

	// 对任务按开始时间进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].start < tasks[j].start
	})

	// 创建一个优先队列（最小堆）来管理任务的结束时间
	pq := &MinHeap{}
	heap.Init(pq)
	result := 0

	for _, t := range tasks {
		heap.Push(pq, t.end)
		// 清除所有不能在当前任务开始时间之前完成的任务
		for pq.Len() > 0 && (*pq)[0] < t.start {
			heap.Pop(pq)
		}

		// 如果有任务可以在当前时间进行，执行它
		if pq.Len() > 0 {
			heap.Pop(pq)
			result++
		}
	}
	fmt.Println(result)
}

// func main() {
// 	taskProcess()
// }
