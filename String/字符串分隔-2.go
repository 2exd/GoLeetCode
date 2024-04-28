package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringSplit() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := scanner.Text()

	// 找到第一个'-'的位置
	pos := strings.Index(s, "-")
	if pos == -1 { // 如果没有'-'，直接输出
		fmt.Println(s)
		return
	}

	// 第一个子串不变，直接追加
	result := strings.Builder{}
	result.WriteString(s[:pos])

	// 处理剩余的字符串
	rest := s[pos+1:]
	rest = strings.ReplaceAll(rest, "-", "") // 移除所有'-'

	// 按照k分组处理剩余的字符串
	for i := 0; i < len(rest); i += k {
		end := i + k
		if end > len(rest) {
			end = len(rest)
		}
		subString := rest[i:end]

		// 计算大小写字母的数量
		lowerCount, upperCount := 0, 0
		for _, c := range subString {
			if c >= 'a' && c <= 'z' {
				lowerCount++
			}
			if c >= 'A' && c <= 'Z' {
				upperCount++
			}
		}

		// 根据大小写字母数量进行转换
		if lowerCount > upperCount {
			subString = strings.ToLower(subString)
		} else if upperCount > lowerCount {
			subString = strings.ToUpper(subString)
		}

		// 添加到结果中
		result.WriteString("-")
		result.WriteString(subString)
	}

	// 输出最终结果
	fmt.Println(result.String())
}
