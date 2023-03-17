package main

import "strconv"

func monotoneIncreasingDigits(N int) int {
	// convert digit to string, easy to use index
	s := strconv.Itoa(N)
	// convert string to byte slice, easy to update
	ss := []byte(s)
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			// front bigger than back, front--, all subsequent numbers are set to 9
			ss[i-1] -= 1
			for j := i; j < n; j++ {
				// set to 9
				ss[j] = '9'
			}
		}
	}
	// convert string to digit
	res, _ := strconv.Atoi(string(ss))
	return res
}
