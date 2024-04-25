package main

import (
	"fmt"
)

func hj10() {
	var n string
	fmt.Scanf("%s", &n)
	m := make(map[rune]int)
	count := 0
	for _, v := range n {
		if m[v] == 0 {
			m[v] = 1
			count++
		}
	}
	fmt.Println(count)
}
