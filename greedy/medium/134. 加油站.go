package main

import (
	"math"
)

/*134. 加油站*/
func myCanCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	min := math.MaxInt32
	// if gas < cost then return -1
	// rest = gas[i] - cost[i] 一天剩下的油, 累加，如果恒大于0，说明从0节点出发
	// 如果累加值 < 0, 说明从非0节点出发
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		curSum += rest
		if curSum < min {
			min = curSum
		}
	}
	// 情况一
	if curSum < 0 {
		return -1
	}
	// 情况二
	if min >= 0 {
		return 0
	}
	// 情况三: 从后向前，看哪个节点能够把这个负数min填平，就是出发节点
	for i := len(gas) - 1; i >= 0; i-- {
		rest := gas[i] - cost[i]
		min += rest
		if min >= 0 {
			return i
		}
	}
	return -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		// optimization, if curSum<0 means that can't start from this node
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

//
// func main() {
// 	gas := []int{1, 2, 3, 4, 5}
// 	cost := []int{3, 4, 5, 1, 2}
// 	negations := canCompleteCircuit(gas, cost)
// 	fmt.Println(negations)
// }
