package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		strs := strings.Fields(sc.Text())
		m, _ := strconv.Atoi(strs[0])
		n, _ := strconv.Atoi(strs[1])
		fmt.Println(suShu(m, n))
	}
}

func suShu(m, n int) int {
	res := 0
	for i := m + 1; i < n; i++ {
		if isSuShu(i) {
			fmt.Println("sushu: ", i)
			res++
		}
	}
	return res
}

func isSuShu(n int) bool {
	if n == 2 {
		return true
	}

	if n > 2 && n%2 == 0 {
		return false
	}
	limit := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
