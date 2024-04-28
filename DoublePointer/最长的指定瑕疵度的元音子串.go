package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	yuanYin()
// }

func yuanYin() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var flaw int
	fmt.Sscanf(scanner.Text(), "%d", &flaw)
	scanner.Scan()
	s := scanner.Text()

	// 元音字符集合
	vowels := "aeiouAEIOU"
	isVowel := make(map[rune]bool)
	for _, v := range vowels {
		isVowel[v] = true
	}

	n := len(s)
	// 为了方便计算瑕疵度，我们创建一个数组，其中flawCnt[i]表示字符串前i个字符的瑕疵度
	flawCnt := make([]int, n+1)
	vowelIndices := make([]int, 0) // 存储元音字符的下标

	for i := 0; i < n; i++ {
		flawCnt[i+1] = flawCnt[i]
		if strings.ContainsRune(vowels, rune(s[i])) {
			// 如果是元音字符，记录下标
			vowelIndices = append(vowelIndices, i)
		} else {
			// 如果是非元音字符，瑕疵度加1
			flawCnt[i+1]++
		}
	}

	maxLength := 0
	// 使用滑动窗口找到瑕疵度符合要求的最长子串
	l := 0
	for r := 0; r < len(vowelIndices); r++ {
		right := vowelIndices[r]
		for l <= r {
			left := vowelIndices[l]
			// 计算当前子串的瑕疵度
			currentFlaw := flawCnt[right+1] - flawCnt[left]
			if currentFlaw > flaw {
				// 如果瑕疵度大于预期，缩小窗口
				l++
			} else if currentFlaw == flaw {
				// 如果瑕疵度符合预期，尝试更新最长长度
				maxLength = max(maxLength, right-left+1)
				break
			} else {
				// 如果当前瑕疵度小于预期，停止当前循环，等待右指针增加后再次检查
				break
			}
		}
	}

	fmt.Println(maxLength)
}

// 辅助函数，用于比较并返回两个整数中的最大值
func maxYuanYin(a, b int) int {
	if a > b {
		return a
	}
	return b
}
