package main

/*1672. 最富有客户的资产总量*/

func maximumWealth(accounts [][]int) (ans int) {
	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if sum > ans {
			ans = sum
		}
	}
	return
}
