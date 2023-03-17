package main

import (
	"fmt"
	"math"
)

/*479. 最大回文数乘积*/
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	upper := int(math.Pow10(n)) - 1
	for left := upper; ; left-- { // 枚举回文数的左半部分
		p := left
		for x := left; x > 0; x /= 10 {
			p = p*10 + x%10 // 翻转左半部分到其自身末尾，构造回文数 p
		}
		for x := upper; x*x >= p; x-- {
			if p%x == 0 { // x 是 p 的因子
				return p % 1337
			}
		}
	}
}

/*复杂度分析

时间复杂度(10^2n)枚举 left 和 x 的时间复杂度均为 O(10^n)。实际上我们只需要枚举远小于 10^n10n个的 left 就能找到答案，实际的时间复杂度远低于 O(10^{2n})

空间复杂度：O(1)，我们只需要常数的空间保存若干变量。*/

func main() {
	fmt.Println(largestPalindrome(3))
}
