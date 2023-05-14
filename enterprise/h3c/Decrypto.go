package main

import (
	"strconv"
	"strings"
)

/*
	模的逆运算
*/
func decryptString(source, key string) string {
	// write code here
	lenS := len(source)
	lenK := len(key)
	sb := new(strings.Builder)
	for i := 0; i < lenS; i++ {
		l := int(source[i] - '0')
		r := int(key[i%lenK] - '0')
		cur := 0
		if l >= r {
			cur = (l - r) % 10
		} else {
			cur = (l - r + 10) % 10
		}
		sb.WriteString(strconv.Itoa(cur))
	}
	return sb.String()
}

// func main() {
// 	s := decryptString("24657980", "123")
// 	fmt.Print(s)
// }
