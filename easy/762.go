package main

import (
	"math"
)

/*762. 二进制表示中质数个计算置位*/
func countPrimeSetBits(left int, right int) int {
	var result = 0
	var oneNum = 0
	for i := left; i <= right; i++ {
		oneNum = 0
		var j = i
		for j > 0 {
			if j&0x01 == 1 {
				oneNum++
			}
			j >>= 1
		}

		if isPrime(oneNum) == false || oneNum == 1 {
		} else {
			result++
		}
	}
	return result
}

func isPrime(num int) bool {
	temp := math.Sqrt(float64(num))
	for i := 2; i <= int(temp); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

/*
func countPrimeSetBits(left int, right int) (ans int) {
	for ; left <= right; left++ {
		if ((1 << bits.OnesCount(uint(left))) & 665772) != 0 {
			ans++
		}
	}
	return
}
*/

//func main(){
//	fmt.Println(countPrimeSetBits(6,10))
//
//	fmt.Println(countPrimeSetBits(10,15))
//}
