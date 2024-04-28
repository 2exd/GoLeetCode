package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func memFlag() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// 读取输入的第一行，表示访存序列的记录条数
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 使用 map 记录每个内存页框号的访问次数
	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		page, _ := strconv.Atoi(scanner.Text())
		counts[page]++
	}

	// 读取热内存的频次阈值
	scanner.Scan()
	threshold, _ := strconv.Atoi(scanner.Text())

	// 过滤并收集访问次数大于等于阈值的内存页框号
	var hotPages []int
	for page, count := range counts {
		if count >= threshold {
			hotPages = append(hotPages, page)
		}
	}

	// 根据题目要求排序：频次高的在前，频次相同则页框号小的在前
	sort.Slice(hotPages, func(i, j int) bool {
		if counts[hotPages[i]] == counts[hotPages[j]] {
			return hotPages[i] < hotPages[j]
		}
		return counts[hotPages[i]] > counts[hotPages[j]]
	})

	// 输出结果
	fmt.Println(len(hotPages))
	for _, page := range hotPages {
		fmt.Println(page)
	}
}

// func main() {
// 	memFlag()
// }
