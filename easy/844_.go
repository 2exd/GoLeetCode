package main

// 栈方法
/*
	遇到#就退一格
	但是此解法时间复杂度和空间复杂度较高，

	时间复杂度：O(N+M)
	空间复杂度：O(N+M)
*/
//func backspaceCompare(s string, t string) bool {
//	return build(s) == build(t)
//}
func build(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) != 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}

/*双指针
遇到# skip++
遇到普通字符，若skip不为0，当前字符删去
			若skip为0，当前字符不删去，比较字符，如果字符不相等就返回false
*/
func backspaceCompare(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}

//
//func main() {
//	var s = "ab#c"
//	var t = "ab#c"
//	backspaceCompare(s, t)
//}
