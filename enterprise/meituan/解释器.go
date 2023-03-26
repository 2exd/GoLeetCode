package main

import (
	"fmt"
	"strings"
)

/*
给一个字符串，以分号分割。
每个key=value当做一个键值对存到字典里。
给一串字符串查找，如果没有这个key输出empty，如果有一个就输出value，如果有多个就输出最新的（可以直接覆盖，这样找到的就是最新的）。
唯一要注意的就是直接按照“;”分割会导致最后多一个空字符串，删了就可以
*/
func console(strOrigin string, intputs []string) {
	m := make(map[string]string, 0)
	strOrigin = strOrigin[:len(strOrigin)-1]
	strs := strings.Split(strOrigin, ";")
	for _, str := range strs {
		kv := strings.Split(str, "=")
		m[kv[0]] = kv[1]
	}

	for _, input := range intputs {
		if _, ok := m[input]; !ok {
			fmt.Println("EMPTY")
		} else {
			fmt.Println(m[input])

		}
	}
}
