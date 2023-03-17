package main

import (
	"strconv"
	"strings"
)

var (
	path93 []string // 放已经回文的子串
	res93  []string
)

func restoreIpAddresses(s string) []string {
	res93 = make([]string, 0)
	path93 = make([]string, 0, len(s))
	dfs93(s, 0)
	return res93
}

func dfs93(s string, start int) {
	if len(path93) == 4 {
		if start == len(s) {
			str := strings.Join(path93, ".")
			res93 = append(res93, str)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' { // 含有前导 0，无效
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path93 = append(path93, str) // 符合条件的就进入下一层
			dfs93(s, i+1)
			path93 = path93[:len(path93)-1]
		} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
			break
		}
	}

}
