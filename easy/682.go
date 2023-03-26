package main

import "strconv"

/*682. 棒球比赛*/

/*方法一：模拟
利用变长数组对栈进行模拟*/
func calPoints(ops []string) (ans int) {
	// 计算得分
	points := []int{}
	for _, op := range ops {
		n := len(points)
		switch op[0] {
		case '+':
			ans += points[n-1] + points[n-2] // points中前两项相加
			// 加入points
			points = append(points, points[n-1]+points[n-2])
		case 'D':
			ans += points[n-1] * 2 // points中前一项*2
			points = append(points, 2*points[n-1])
		case 'C':
			ans -= points[n-1] // 减去前一项的值
			points = points[:len(points)-1]
		default:
			pt, _ := strconv.Atoi(op) // 字符串转int
			ans += pt
			points = append(points, pt)
		}
	}
	return
}

/*复杂度分析
时间复杂度：O(n)，其中 n 为数组ops 的大小。遍历整个 ops 需要 O(n)。
空间复杂度：O(n)。变长数组最多保存 O(n)个元素。*/
