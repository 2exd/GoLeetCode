package main

import (
	"fmt"
	"math"
)

func chaoChe() {
	var m, n int
	fmt.Scan(&m, &n)

	maxTime := 0.0
	for i := 0; i < m; i++ {
		var v int
		fmt.Scan(&v)
		currTime := float64(n)/float64(v) + float64(i)
		maxTime = math.Max(maxTime, currTime)
	}
	fmt.Println(maxTime - float64(m-1))
}

func main() {
	chaoChe()
}
