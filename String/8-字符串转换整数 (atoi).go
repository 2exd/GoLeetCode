package main

import (
	"math"
)

/*8-字符串转换整数 (atoi)*/
func myAtoi(s string) int {
	num, sign, i, n := 0, 1, 0, len(s)
	// step 1: drop space ''
	for i < n && s[i] == ' ' {
		i++
	}

	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	for i < n && s[i] <= '9' && s[i] >= '0' {
		num = num*10 + int(s[i]-'0')
		if num > math.MaxInt32 {
			if sign == -1 {
				return sign*math.MaxInt32 - 1
			}
			return math.MaxInt32
		}
		i++
	}
	return sign * num
}

// func main() {
// 	fmt.Print(myAtoi("-91283472332"))
// }
