package main

/*728. 自除数*/
func myselfDividingNumbers(left int, right int) []int {
	var list []int
	for i := left; i <= right; i++ {
		flag := true
		for j := i; j > 0; {
			k := j % 10
			if k == 0 {
				flag = false
				break
			}
			if i%k != 0 {
				flag = false
			}
			j /= 10
		}
		if flag {
			list = append(list, i)
		}
	}
	return list
}

/*官方*/
func isSelfDividing(num int) bool {
	for x := num; x > 0; x /= 10 {
		if d := x % 10; d == 0 || num%d != 0 {
			return false
		}
	}
	return true
}

func selfDividingNumbers(left, right int) (ans []int) {
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return
}

/*复杂度分析
时间复杂度：O(nlogright)，其中 n 是范围内的整数个数，right 是范围内的最大整数。对于范围内的每个整数，需要 O(logright) 的时间判断是否为自除数。
空间复杂度：O(1)。除了返回值以外，使用的额外空间为 O(1)。*/

// func main() {
//	//fmt.Println(selfDividingNumbers(1, 22))
//	fmt.Println(selfDividingNumbers(47, 85))
// }
