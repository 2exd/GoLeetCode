package main

/*3. 无重复字符的最长子串*/

func max3(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	cnt := make(map[byte]int) // 也可以用 map[byte]int，这里为了方便用的数组
	for right, c := range s {
		cnt[byte(c)]++
		for cnt[byte(c)] > 1 { // 不满足要求
			cnt[s[left]]--
			left++
		}
		ans = max3(ans, right-left+1)
	}
	return
}
