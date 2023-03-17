package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*1363. 形成三的最大倍数*/
/*思路
先把所有数加入到hash表中
计算总的数加对3的模
总和为0，直接返回“0”
①模为0，直接返回计算结果返回
②模为1，且存在余为1的项，删除最小的余为1的项，计算结果
③模为1，不存在余为1的项，删除2个最小的余为2的项，计算结果
④模为2，且存在余为2的项，删除最小的余为2的项，计算结果
⑤模为2，不存在余为2的项，删除2个最小的余为1的项，计算结果
返回最终结果*/
func mylargestMultipleOfThree(digits []int) string {
	// 降序排列
	sort.Ints(digits)
	sum := 0
	var m1 []int
	var m2 []int

	// 计算总和
	for i := range digits {
		sum += digits[i]
		if digits[i]%3 == 1 {
			m1 = append(m1, digits[i])
		}
		if digits[i]%3 == 2 {
			m2 = append(m2, digits[i])
		}
	}

	if sum == 0 {
		return "0"
	}

	var result string
	// 如果是三的倍数就计算result
	if sum%3 == 0 {
		result = helper(digits)
	} else if sum%3 == 1 && len(m1) > 0 { // 总和余为1且存在余1的数，删除一个最小的余为1的数
		// 删除一个最小的余为1的数
		for i := 0; i < len(digits); i++ {
			if digits[i]%3 == 1 {
				digits[i] = -1
				break
			}
		}
		result = helper(digits)

		if result == "" {
			return result
		}
		// 如果result = "000000" 就返回"0"
		test, _ := strconv.Atoi(result)
		if test == 0 {
			return "0"
		}
		return result
	} else if sum%3 == 1 && len(m1) == 0 { // 总和余为1且不存在余1的数，删除两个最小的余为2的数
		count := 0
		for i := 0; i < len(digits); i++ {
			if digits[i]%3 == 2 {
				digits[i] = -1
				count++
				// 删除两个就break
				if count == 2 {
					break
				}
			}
		}
		result = helper(digits)
		if result == "" {
			return result
		}
		// 如果result = "000000" 就返回"0"
		test, _ := strconv.Atoi(result)
		if test == 0 {
			return "0"
		}
		return result
	} else if sum%3 == 2 && len(m2) > 0 { // 总和余为2且存在余2的数，删除一个最小的余为2的数
		// 删除一个最小的余为2的数
		for i := 0; i < len(digits); i++ {
			if digits[i]%3 == 2 {
				digits[i] = -1
				break
			}
		}
		result = helper(digits)
		// 如果result = "000000" 就返回"0"
		if result == "" {
			return result
		}
		test, _ := strconv.Atoi(result)
		if test == 0 {
			return "0"
		}
		return result
	} else if sum%3 == 2 && len(m2) == 0 { // 总和余为2且不存在余2的数，删除两个最小的余为1的数
		count := 0
		for i := 0; i < len(digits); i++ {
			if digits[i]%3 == 1 {
				digits[i] = -1
				count++
				// 删除两个就break
				if count == 2 {
					break
				}
			}
		}
		result = helper(digits)
		if result == "" {
			return result
		}
		// 如果result = "000000" 就返回"0"
		test, _ := strconv.Atoi(result)
		if test == 0 {
			return "0"
		}
		return result
	}

	return result
}

// 思路1 ： 之和的 变形题
// 不处理值为-1的项
func helper(nums []int) string {
	var result string
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != -1 {
			result += strconv.Itoa(nums[i])
		}
	}

	return result
}

/*第二名*/
/*分为余为1和余为2的数组m,n,并降序排列
m,n 1:1放入ans数组中(重点：不能全部放完)
再三个三个单独放入，保证放入ans数组中总和为0
m,n剩下的元素再放入1:1放入ans数组中
计算结果并返回
*/
func largestMultipleOfThree(digits []int) string {
	count := len(digits)
	anses := make([]byte,0)
	one := make([]byte,0)
	two := make([]byte,0)
	for i:=0;i<count;i++ {
		if digits[i] % 3 == 0 {
			anses = append(anses,byte(digits[i]+int('0')))
		}else if digits[i] % 3 == 1 {
			one = append(one,byte(digits[i]+int('0')))
		}else if digits[i] % 3 == 2 {
			two = append(two,byte(digits[i]+int('0')))
		}
	}
	// 降序排列
	sort.Slice(one, func(i, j int) bool {
		s, t := one[i], one[j]
		return s<t
	})
	sort.Slice(two, func(i, j int) bool {
		s, t := two[i], two[j]
		return s<t
	})
	len1 := len(one)
	len2 := len(two)
	i := len1-1
	j := len2-1
	// 余为1和余为2的1:1放入,保证一方>0
	for i>0 && j>0 {
		anses = append(anses,one[i])
		i--
		anses = append(anses,two[j])
		j--
	}
	// 余为1的三个三个放
	for i>1 {
		anses = append(anses,one[i])
		i--
		anses = append(anses,one[i])
		i--
		anses = append(anses,one[i])
		i--
	}
	// 余为2的三个三个放
	for j>1 {
		anses = append(anses,two[j])
		j--
		anses = append(anses,two[j])
		j--
		anses = append(anses,two[j])
		j--
	}
	//for i>=0 && j>=0 {
	//	anses = append(anses,one[i])
	//	i--
	//	anses = append(anses,two[j])
	//	j--
	//}

	/*
	   var do func(){
	       if i>0 && j>0 {
	           anses = append(anses,one[i])
	           i--
	           anses = append(anses,one[i])
	           i--
	           anses = append(anses,two[j])
	           j--
	           anses = append(anses,two[j])
	           j--
	       }else if
	           do()
	   }
	*/

	if len(anses) < 1 {
		return ""
	}
	sort.Slice(anses, func(i, j int) bool {
		s, t := anses[i], anses[j]
		return s>t
	})

	if anses[0] == '0' {
		return "0"
	}
	return string(anses)
}

func main() {
	//arr := []int{8, 6, 7, 1, 0, 0, 8, 9, 7, 5, 3, 6, 3, 2}
	//arr := []int{8,6,7,1,0}
	//arr := []int{5, 8}
	//arr := []int{1}
	//arr := []int{0}
	arr := []int{0,0,0,0,0,0,1}
	//arr := []int{8, 1, 9}
	fmt.Println(largestMultipleOfThree(arr))

}
