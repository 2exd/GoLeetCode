package main

/*5-最长回文子串*/
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n) // dp[i][j]=true表示s[i,j]左闭右闭区间能构成回文子串

	begin := 0
	maxLen := 1 // 因为题意说至少s.len >= 1
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	} // len = 1

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			begin = i
			maxLen = 2
		}
	} // len = 2
	for j := 2; j < n; j++ { // len >= 3
		for i := 0; i <= j; i++ { // [i,j]左闭右闭区间
			if j-i < 2 {
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] { // 找到一个回文子串了
					len_ := j - i + 1
					if len_ > maxLen {
						maxLen = len_
						begin = i
					}
				}
			}
		}
	}

	return s[begin : begin+maxLen]
}
