package main

/*202. 快乐数 easy*/

func isHappy(n int) bool {
	m := make(map[int]bool)

	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}
