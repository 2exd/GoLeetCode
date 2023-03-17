package main

/*LCP 17. 速算机器人*/

/*方法一 遍历
*/
func calculate(s string) int {
	x := 1
	y := 0
	//var result int
	for i:=0; i<len(s);i++ {
		switch s[i] {
		case 'A':
			x = 2*x + y
		case 'B':
			y = 2*y + x
		}
	}

	return x+y
}
/*复杂度分析
时间复杂度：O(n)，其中 n 为String s的大小。遍历整个 s 需要 O(n)。
空间复杂度：O(1)。*/

/*方法二 数学法
分析发现
出现一个"A"，有x+y=(2x+y)+y=2x+2y
出现一个"B"，有x+y=x+(2y+x)=2x+2y
所以每出现一个A/B，都使x+y的值翻倍
因此结果是2*len(s)
*/
func calculate2(s string) int {

	return 1 << len(s)
}
/*复杂度分析
时间复杂度：O(1)。
空间复杂度：O(1)。*/

