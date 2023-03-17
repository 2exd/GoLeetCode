package main

import "sort"

type pair struct {
	target  string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findItinerary(tickets [][]string) []string {

	res332 := make([]string, 0, len(tickets)+1)
	targets := make(map[string]pairs, len(tickets))

	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}

		targets[ticket[0]] = append(targets[ticket[0]], &pair{ticket[1], false})
	}

	for k, _ := range targets {
		sort.Sort(targets[k])
	}

	res332 = append(res332, "JFK")

	var backtracking func() bool

	backtracking = func() bool {
		// 结束条件, 机票全部使用
		if len(res332) == len(tickets)+1 {
			return true
		}

		// 每次从上一次目的地出发
		for _, pair := range targets[res332[len(res332)-1]] {
			// 未被访问过
			if pair.visited == false {
				res332 = append(res332, pair.target)
				pair.visited = true

				if backtracking() {
					return true
				}

				res332 = res332[:len(res332)-1]
				pair.visited = false
			}
		}

		return false
	}

	backtracking()

	return res332
}
