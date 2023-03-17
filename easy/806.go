package main

/*806. 写字符串需要的行数*/
func numberOfLines(widths []int, s string) []int {
	ans := make([]int, 2)
	var sum int
	for i := 0; i < len(s); i++ {
		if sum+widths[s[i]-'a'] <= 100 {
			sum += widths[s[i]-'a']
		} else {
			ans[0]++
			sum = widths[s[i]-'a']
		}
	}
	ans[1] = sum
	if ans[1] != 0 {
		ans[0]++
	}
	return ans
}
