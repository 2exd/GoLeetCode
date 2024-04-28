package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func passwordDecrypt() {
	// 创建一个 Scanner 从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	encryptedText := scanner.Text()

	// 存储解密后的文本
	var decryptedText strings.Builder

	// 从后向前遍历加密的字符串，这样可以方便处理带星号的多位数
	for i := len(encryptedText) - 1; i >= 0; {
		if encryptedText[i] == '*' {
			// 处理两位数加星号的情况（如10*到26*）
			if num, err := strconv.Atoi(encryptedText[i-2 : i]); err == nil {
				// 将数字转换为对应的字符
				decryptedText.WriteByte(byte('j' + num - 10))
				i -= 3 // 移动到下一个数字或星号的位置
			}
		} else {
			// 处理一位数的情况（如1到9）
			decryptedText.WriteByte(byte('a' + encryptedText[i] - '1'))
			i-- // 移动到下一个数字的位置
		}
	}

	// 由于是从后向前解析，所以需要反转解密后的字符串
	// 输出解密后的文本
	fmt.Println(reverseStringP(decryptedText.String()))
}

// reverseString 反转字符串
func reverseStringP(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
