package main

/*2024. 考试的最大困扰度*/
/*典型的滑动窗口题*/
func mymaxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			// 计算不为ch的总数
			sum++
		}
		// 不为ch的总数大于k,收缩窗口，简言之窗口内不为ch的字符数必须小于k
		for sum > k {
			// 如果窗口左边不为ch,sum-1
			if answerKey[left] != ch {
				sum--
			}
			// 左边界收缩
			left++
		}
		// 保存历史窗口最大值
		ans = max(ans, right-left+1)
	}
	return
}

/*更高效*/
func maxConsecutiveAnswers(answerKey string, k int) int {
	// 保存窗口中A-Z的字符数
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range answerKey {
		cnt[ch-'A']++
		// 某字符的最大值
		// 每次都更新，不用像第一种做法遍历两次
		maxCnt = max(maxCnt, cnt[ch-'A'])
		// 大于k，左边界收缩
		if right-left+1-maxCnt > k {
			cnt[answerKey[left]-'A']--
			left++
		}
	}
	// maxCnt = max(maxCnt, cnt[ch-'A']) 保证了len(answerKey) - left是最大值
	return len(answerKey) - left
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/*复杂度分析
时间复杂度：O(n)，其中 n 为字符串长度，我们只需要遍历该字符串两次。
空间复杂度：O(1)，我们只需要常数的空间保存若干变量。。*/

//func main()  {
//
//	var a byte = 'T'
//	fmt.Printf("%d %T\n", a, a)
//
//	var b rune='F'
//	fmt.Printf("%d %T\n", b, b)
//}
