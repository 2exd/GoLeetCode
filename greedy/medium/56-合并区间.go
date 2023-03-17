package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 	first order by start position asc
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// merge the coincident intervals
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			// merge
			intervals[i][1] = max56(intervals[i][1], intervals[i+1][1])
			// remove middle interval
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			// still judge from the current interval
			i--
		}
	}
	return intervals
}
func max56(a, b int) int {
	if a < b {
		return b
	}
	return a
}
