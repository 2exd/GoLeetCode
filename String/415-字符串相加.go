package main

import (
	"strings"
)

/*415-字符串相加*/
func addStrings(num1 string, num2 string) string {
	i1, i2 := len(num1)-1, len(num2)-1
	var carry byte
	var sb strings.Builder
	for i1 >= 0 || i2 >= 0 {
		var n1, n2 byte
		if i1 >= 0 {
			n1 = num1[i1] - '0'
		}
		if i2 >= 0 {
			n2 = num2[i2] - '0'
		}
		x := n1 + n2 + carry
		sb.WriteByte((x % 10) + '0') // get digit
		carry = x / 10               // get carry
		i1--
		i2--
	}
	// get carry
	if carry > 0 {
		sb.WriteByte(carry + '0')
	}
	// reverse string
	res := []rune(sb.String())
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
