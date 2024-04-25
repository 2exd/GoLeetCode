package main

import (
	"bufio"
	"fmt"
	"os"
)

func hj20() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		s := reader.Text()
		if valid(s) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

func valid(s string) bool {
	// 1. 长度大于8
	if len(s) <= 8 {
		return false
	}
	// 2. 先统计每种字符的个数，看看是不是至少3个大于0
	counts := [4]int{0}
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			counts[0]++
		} else if v >= 'A' && v <= 'Z' {
			counts[1]++
		} else if v >= '0' && v <= '9' {
			counts[2]++
		} else {
			counts[3]++
		}
	}
	diffType := 0
	for _, v := range counts {
		if v != 0 {
			diffType++
		}
	}
	if diffType < 3 {
		return false
	}
	// 3. 取大小为3的滑动窗口，hash解之
	m := make(map[string]bool)
	for i := 0; i < len(s)-2; i++ {
		if _, ok := m[s[i:i+3]]; !ok {
			m[s[i:i+3]] = true
		} else {
			return false
		}
	}

	return true
}
