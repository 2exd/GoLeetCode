package main

/*172. 阶乘后的零*/
func trailingZeroes(n int) (ans int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return
}

/*时间复杂度：O(n)。
空间复杂度：O(1)。*/

// 不断除以5
func trailingZeroes2(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

/*时间复杂度：O(logn)。
空间复杂度：O(1)。*/
