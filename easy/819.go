package main

import (
	"unicode"
)

/*819. 最常见的单词*/

func mostCommonWord(paragraph string, banned []string) string {
	// 记录被ban的单词
	ban := map[string]bool{}
	for _, s := range banned {
		ban[s] = true
	}
	freq := map[string]int{}
	maxFreq := 0
	var word []byte
	for i, n := 0, len(paragraph); i <= n; i++ {
		// 不是字符就停止
		if i < n && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			s := string(word)
			if !ban[s] {
				freq[s]++
				maxFreq = max(maxFreq, freq[s])
			}
			word = nil
		}
	}
	for s, f := range freq {
		if f == maxFreq {
			return s
		}
	}
	return ""
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/*
复杂度分析

时间复杂度：O(n+m)，其中 n 是段落 paragraph 的长度，m 是禁用单词列表 banned 的长度。遍历禁用单词列表一次将禁用单词存入哈希集合中需要 O(m) 的时间，
遍历段落得到每个非禁用单词的计数需要 O(n) 的时间，遍历哈希表得到最常见的单词需要 O(n) 的时间。

空间复杂度：O(n+m)，其中 n 是段落 paragraph 的长度，m 是禁用单词列表 banned 的长度。存储禁用单词的哈希集合需要 O(m) 的空间，记录每个单词的计数的哈希表需要 O(n) 的空间。
*/
