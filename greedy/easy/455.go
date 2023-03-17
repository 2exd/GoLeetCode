package main

import "sort"

/*455. 分发饼干*/

func findContentChildren(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(size), len(greed)
	i := 0
	j := 0
	for i < n && j < m {
		if size[i] >= greed[j] {
			j++
		}
		i++
	}
	return j
}
