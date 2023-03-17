package main

import "fmt"

/*904. 水果成篮*/
func totalFruit(fruits []int) int {
	mapVariable := make(map[int]int)
	i := 0
	sum := 0
	l := len(fruits) // 数组长度
	result := 0      // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		mapVariable[fruits[j]]++
		sum++
		for len(mapVariable) > 2 {
			mapVariable[fruits[i]]--
			sum--
			if mapVariable[fruits[i]] == 0 {
				delete(mapVariable, fruits[i])
			}
			i++
		}
		if sum > result {
			result = sum
		}

	}

	if result == l+1 {
		return 0
	} else {
		fmt.Print(result)
		return result
	}
}

//func main() {
//	nums := []int{0, 1, 2, 2}
//	totalFruit(nums)
//}
