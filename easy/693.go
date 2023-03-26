package main

/*693. 交替位二进制数*/
/*
变量i = 当前最末位的值
n每次右移一位
i 与 移位后的n的最末位比较 是否相同
*/
func myhasAlternatingBits(n int) bool {
	i := n & 1
	for {
		n >>= 1
		if 1&n == i {
			return false
		}
		i = 1 & n
		if n == 0 {
			break
		}
	}
	return true
}

/*精简版*/
/*按位补码(NOT) 异或操作
如果为交替数，n ^ n>>1 为全1
那么 n & n+1 结果应该为0
*/
func hasAlternatingBits(n int) bool {
	a := n ^ n>>1
	return a&(a+1) == 0
}

/*复杂度分析
时间复杂度：O(1)。仅使用了常数时间来计算。
空间复杂度：O(1)。使用了常数空间来存储中间变量*/

// func main()  {
//	fmt.Println(hasAlternatingBits(5))
// }
