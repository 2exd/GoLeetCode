package main

import "fmt"

/*
现在小美获得了一个字符串。小美想要使得这个字符串是回文串。
小美找到了你。你可以将字符串中至多两个位置改为任意小写英文字符 'a'-'z'
你的任务是帮助小美在当前制约下，获得字典序最小的回文字符串。
*/
func main() {
	huiwen()
}

func huiwen() {
	var str string
	fmt.Scanf("%s\n", &str)
	n := len(str)
	s := []byte(str)
	l, r := 0, n-1
	cnt := 0

	// 最多修改两次
	for l < r {
		if cnt == 2 {
			break
		}
		if s[l] != s[r] {
			s[l] = 'a'
			s[r] = 'a'
			cnt++
		}
		l++
		r--
	}

	for cnt != 2 {
		if l >= r {
			l, r = 0, n-1
			for l < r {
				if cnt == 2 {
					break
				}
				if s[l] != 'a' {
					s[l] = 'a'
					s[r] = 'a'
					cnt++
				}
				l++
				r--
			}
			if l == r && cnt < 2 {
				s[l] = 'a'
				l++
				r--
			}
		}
	}
	fmt.Println(s)
}
