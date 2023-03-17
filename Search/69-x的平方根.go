package main

// 精确到个位
func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := (l + r) >> 1
		if x < mid*mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

// 精确到3位小数
func mySqrt3(x float64) float64 {
	l, r := 0.0, x
	for l <= r {
		mid := (l + r) / 2
		if x < mid*mid {
			r = mid - 1e-3
		} else {
			l = mid + 1e-3
		}
	}
	return r
}
