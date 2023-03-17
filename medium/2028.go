package main

/*2028.找出缺失的观测数据*/

func mymissingRolls(rolls []int, mean int, n int) []int {
	totalN := mean * (n + len(rolls))
	var rollsN = make([]int, n)
	for _, value := range rolls {
		totalN -= value
	}

	// 平均点数,余数
	x, y := totalN/n, totalN%n

	if x > 6 || x <= 0 {
		return nil
	}
	if x == 6 && y > 0 {
		return nil
	}

	for i := 0; i < n; i++ {
		if y > 0 {
			rollsN[i] = x + 1
			y--
		} else {
			rollsN[i] = x
		}
	}

	return rollsN
}

func missingRolls(rolls []int, mean, n int) []int {
	missingSum := mean * (n + len(rolls))
	for _, roll := range rolls {
		missingSum -= roll
	}
	if missingSum < n || missingSum > n*6 {
		return nil
	}

	quotient, remainder := missingSum/n, missingSum%n
	ans := make([]int, n)
	for i := range ans {
		ans[i] = quotient
		if i < remainder {
			ans[i]++
		}
	}
	return ans
}

//func main() {
//	var rolls = []int{3, 5, 3}
//	fmt.Println(missingRolls(rolls, 5, 3))
//}
