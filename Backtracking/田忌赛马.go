package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	tianJiSaiMa()
// }

func tianJiSaiMa() {
	// 全排列，回溯问题
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	a := parseInput(reader.Text())
	reader.Scan()
	b := parseInput(reader.Text())

	solution := NewSolution(a, b)
	solution.Solve(0, 0)
	fmt.Println(solution.ResultCnt)
}

func parseInput(input string) []int {
	strs := strings.Split(input, " ")
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

type Solution struct {
	a, b      []int
	n         int
	MaxCnt    int
	ResultCnt int
	vis       []bool
}

func NewSolution(a, b []int) *Solution {
	return &Solution{
		a:         a,
		b:         b,
		n:         len(a),
		MaxCnt:    0,
		ResultCnt: 0,
		vis:       make([]bool, len(a)),
	}
}

func (s *Solution) Solve(idx, cnt int) {
	// 剪枝：如果剩余可能性加上当前计数小于已知的最大计数，终止这个分支
	if s.n-idx+cnt < s.MaxCnt {
		return
	}
	// 如果所有元素都已处理，更新最优解计数
	if idx == s.n {
		if cnt > s.MaxCnt {
			s.MaxCnt = cnt
			s.ResultCnt = 1
		} else if cnt == s.MaxCnt {
			s.ResultCnt++
		}
		return
	}

	// 遍历每个元素，如果未被访问，进行递归
	for i := 0; i < s.n; i++ {
		if s.vis[i] {
			continue
		}
		s.vis[i] = true
		s.Solve(idx+1, cnt+boolToInt(s.a[i] > s.b[idx]))
		s.vis[i] = false
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
