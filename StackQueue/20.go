package main

/*20. 有效的括号 easy*/

func isValid(s string) bool {
	if s == "" {
		return true
	}
	n := len(s)
	if n%2 == 1 {
		return false
	}
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)

	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			stack = append(stack, byte(v))
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[byte(v)] {
			// 后进先出
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
