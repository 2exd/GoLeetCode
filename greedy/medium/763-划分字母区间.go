package main

/*763. 划分字母区间*/
func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s), 0, 0
	for i := 0; i < size; i++ {
		// find the longest position for each character
		marks[s[i]-'a'] = i
	}
	for i := 0; i < size; i++ {
		right = max763(right, marks[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}

func max763(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
