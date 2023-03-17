package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*1606. 找到处理最多请求的服务器*/

type server struct {
	taskNum     int
	taskEndTime int
}

func creatServer() server {
	return server{0, 0}
}

func maxTask(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func mybusiestServers(k int, arrival []int, load []int) []int {
	var servers []server
	var maxTaskNum int
	var busyServers []int
	// 初始化
	for i := 0; i < k; i++ {
		servers = append(servers, creatServer())
	}

	for index, time := range arrival {
		for j := 0; j < k; j++ {
			s := (index + j) % k
			if s == k {
				s = 0
			}
			// 该服务器繁忙,进入下一次循环
			if servers[s].taskEndTime > time {
			} else { // 该服务器空闲,获取该任务
				// 任务数量+1
				servers[s].taskNum++
				// 任务完成时间 = 任务所需时间 + 开始任务时间
				servers[s].taskEndTime = load[index] + time
				// 更新最大任务数量
				maxTaskNum = maxTask(maxTaskNum, servers[s].taskNum)
				break
			}
		}
	}
	for i := 0; i < k; i++ {
		if servers[i].taskNum == maxTaskNum {
			busyServers = append(busyServers, i)
		}
	}
	return busyServers
}

/*时间复杂度：O(k * n + 2K)，其中 k 为服务器的数目
空间复杂度：O(k)。
超出时间限制*/

/*双优先队列*/
func busiestServers(k int, arrival []int, load []int) []int {
	used, available, cnts := &ListHeap{}, &IntHeap{}, make([]int, k)
	for i := 0; i < k; i++ {
		// 将空闲服务器的编号都放入一个有序集合available
		heap.Push(available, i)
	}
	for i := 0; i < len(arrival); i++ {
		arr, duration := arrival[i], load[i]
		for used.Len() > 0 && (*used)[0][0] <= arr {
			cur := heap.Pop(used).([]int)
			heap.Push(available, i+((cur[1]-i)%k+k)%k)
		}
		if available.Len() > 0 {
			idx := heap.Pop(available).(int) % k
			cnts[idx]++
			heap.Push(used, []int{arr + duration, idx})
		}
	}
	// 找到最大任务数,并确定最繁忙服务器列表
	ans, m := []int{}, 0
	for i := 0; i < k; i++ {
		if cnts[i] > m {
			m = cnts[i]
			ans = []int{i}
		} else if cnts[i] == m {
			ans = append(ans, i)
		}
	}
	return ans
}

type ListHeap [][]int

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

/*复杂度分析
时间复杂度：O((k + n) \log k) 或 O(k + n \log k)，其中 k 为服务器的数目，n 为请求的数目。开始时 available 放入所有的服务器的时间复杂度为 O(k \log k) 或 O(k)，
取决于语言实现；在处理请求时，busy 最多执行 n 次放入和移除操作，available 最多执行 n 次放入和移除操作，因此时间复杂度为 O(n \log k)；获取最繁忙服务器列表的时间复杂度为 O(k)。
空间复杂度：O(k)。busy 和 available 最多放入 kk 个元素，需要 O(k) 的空间；requests 需要 O(k) 的空间。*/

func BbusiestServers(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)}
	// 将空闲服务器的编号都放入一个优先队列available
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}
	// 将正在处理请求的服务器的处理结束时间和编号都放入一个优先队列busy
	busy := hp{}
	// 用一个数组 requests 记录对应服务器处理的请求数目
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			heap.Push(&available, i+((busy[0].id-i)%k+k)%k) // 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Pop(&busy)
		}
		if available.Len() == 0 {
			continue
		}
		id := heap.Pop(&available).(int) % k
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
	}
	return
}

type hi struct{ sort.IntSlice } // 小顶堆

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end } // 小顶堆
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
	// k := 3
	// arrival := []int{1,2,3,4,5}
	// load := []int{5,2,3,3,3}
	// k := 3
	// arrival := []int{1,2,3,4}
	// load := []int{1,2,1,2}
	// k := 3
	// arrival := []int{1,2,3}
	// load := []int{10,12,11}
	// k := 3
	// arrival := []int{1,2,3,4,8,9,10}
	// load := []int{5,2,10,3,1,2,2}
	// k := 1
	// arrival := []int{1}
	// load := []int{1}
	k := 6
	arrival := []int{1, 2, 3, 9, 11, 12, 14}
	load := []int{12, 3, 8, 13, 6, 10, 14}
	fmt.Println(BbusiestServers(k, arrival, load))
}
