package main

import "strings"

/*804. 唯一摩尔斯密码词*/
/*哈希表*/
var morse = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
	"....", "..", ".---", "-.-", ".-..", "--", "-.",
	"---", ".--.", "--.-", ".-.", "...", "-", "..-",
	"...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	set := map[string]struct{}{}
	for _, word := range words {
		trans := &strings.Builder{}
		for _, ch := range word {
			trans.WriteString(morse[ch-'a'])
		}
		// 空结构体，不占用内存
		set[trans.String()] = struct{}{}
	}

	// 只需要返回map元素个数
	return len(set)
}

/*
复杂度分析
时间复杂度：O(S)，其中 S 是数组 words 中所有单词的长度之和。
空间复杂度：O(S)，其中 S 是数组 words 中所有单词的长度之和。

*/
