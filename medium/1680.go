package main

import (
	"math"
	"math/bits"
)

/*1680. 连接连续二进制数字*/
/*拼接
每次拼接后都取模
再计算余*/
func myconcatenatedBinary(n int) int {
	var mod = 1000000007
	var ans int
	for i := 1; i <= n; i++ {
		ans = ans<<bits.Len(uint(i)) | i
		ans %= mod
	}
	return ans
}

/*复杂度分析
时间复杂度：O(n)。
空间复杂度：O(1)。
*/

/*方法二：数学*/
/*
考虑到第i个数字，其二进制位数为bits，i之前的数字的结果为res，那么拼上i之后为：
res左移bits位，加上i
求i的二进制位数：
若i&(i-1)==0，表明i是2的整数次幂，相对于i-1的二进制位数需要+1。
若从n开始倒着遍历，移位操作为(i<<len(bin(res)))+res，由于res二进制位数越积越大，移位成本巨大，导致超时。
*/
func mathconcatenatedBinary(n int) int {
	var mod float64 = 1000000007
	var zeroes float64
	var ans float64
	for k := 0.0; k < 64; k++ {
		lb := math.Pow(2, k-1)
		if lb <= float64(n) {
			t := float64(n) - lb
			u := math.Mod(math.Pow(2, t*k), mod)
			v := math.Mod(math.Pow(math.Pow(2, k)-1, mod-2), mod)
			w := math.Mod(math.Pow(2, (t+1)*k), mod)
			x := math.Mod(math.Pow(2, zeroes), mod)

			ans += math.Mod((math.Pow(2, k)*(u-1)*v+(float64(n)-t)*w-float64(n))*v*x, mod)
			zeroes += (t + 1) * k
			n = int(lb - 1)
		}
	}
	ans += math.Mod(math.Pow(2, zeroes), mod)
	return int(math.Mod(ans, mod))
}

func concatenatedBinary(n int) (res int) {
	mod := int(1e9 + 7)
	// 第i个数字的二进制位数
	bits := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			bits++
		}
		res = ((res << bits) + i) % mod
	}
	return
}

// func main()  {
//	fmt.Println(mathconcatenatedBinary(3))
// }
