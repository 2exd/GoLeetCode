package main

var m map[int]int

// func main() {
// 	m = make(map[int]int)
// 	var n, a, b int
// 	minTime, maxTime, sum := 9999, 0, 0
// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&a, &b)
// 		if b > maxTime {
// 			maxTime = b
// 		}
// 		if a < minTime {
// 			minTime = a
// 		}
// 		for i := a; i <= b; i++ {
// 			m[i]++
// 		}
// 	}
// 	for i := minTime; i <= maxTime; i++ {
// 		sum += getState(i)
// 	}
// 	fmt.Print(sum)
// }

func getState(nowTime int) int {
	taskNum, ok := m[nowTime]
	if !ok {
		return 1
	} else if taskNum == 1 {
		return 3
	} else {
		return 4
	}
}
