package main

import (
	"bufio"
	"fmt"
	"os"
)

func hj23() {

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str := reader.Text()
	res := removeLeast(str)
	fmt.Println(res)
}

func removeLeast(s string) string {
	// 1. 统计每个字符出现的次数
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}
	// 2. 找到出现次数最少的字符
	min := len(s) - 1
	for _, v := range m {
		if v > 0 && v < min {
			min = v
		}
	}
	// 3. 删除所有出现次数最少的字符
	var res []rune
	for _, v := range s {
		if m[v-'a'] > min {
			res = append(res, v)
		}
	}
	return string(res)
}
