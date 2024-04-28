package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMinSort() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	// Convert string to a slice of runes to handle individual characters
	cs := []rune(s)
	lenCs := len(cs)

	// Iterate over the string to find the first beneficial swap
	for left := 0; left < lenCs; left++ {
		// 右侧寻找小于 cs[left] 的最小字符，如果有多个则取最右边的一个
		idx := left
		for right := left + 1; right < lenCs; right++ {
			if cs[right] < cs[idx] {
				idx = right
			}
		}

		// Perform the swap if it results in a smaller string
		if cs[idx] != cs[left] {
			cs[idx], cs[left] = cs[left], cs[idx]
			break
		}
	}

	// Output the potentially modified string
	fmt.Println(string(cs))
}
