package main

/*541. 反转字符串 II easy*/

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	// 每次前进2k步
	for i := 0; i < length; i += 2 * k {
		// 判断是否超出数组范围
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

// 反转
func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
