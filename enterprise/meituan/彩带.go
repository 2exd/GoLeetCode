package main

/*
小美现在有一串彩带，假定每一厘米的彩节上都是一种色彩。
因为任务的需要，小美希望从彩带上截取一段，使得彩带中的颜色数量不超过x种。
显然，这样的截取方法可能非常多，于是小美决定尽最长地截取一段。
你的任务是帮助小美截取尽量长的一段，使得这段彩带上不同的色彩数量不超过K种。
输入描述：
第一行两个整数N，K。以空格分开，分别表示彩带有N厘米长，你截取的一段连续的彩带不能超过k种颜色。
接下来一行N个整数，每个正数表示一种色彩，相同的整数表示相同的色彩。
1<=N，<=500   彩带上的颜色数字介于[1,2000]之间
输出描述：
一行，一个整数，表示选取的彩带的最大长度。
*/
func caidai(n, k int, nums []int) int {
	max := 0
	l, r := 0, 0
	m := make(map[int]int)
	m[nums[0]]++
	cnt := 1
	for i := 1; i < n; i++ {
		if r-l+1 > max {
			max = r - l + 1
		}

		if m[nums[i]] == 0 {
			cnt++
		}
		m[nums[i]]++
		r++
		// 移除
		if cnt > k {
			for l < r {
				m[nums[l]]--
				l++
				if m[nums[l-1]] == 0 {
					cnt--
					break
				}
			}
		}

	}
	return max
}

// func main() {
// 	var num int
// 	var n, k int
//
// 	var nums []int
// 	fmt.Scanf("%d", &n)
// 	fmt.Scanf("%d\n", &k)
// 	for {
// 		_, err := fmt.Scanf("%d", &num)
// 		if err == nil {
// 			nums = append(nums, num)
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Println(caidai(n, k, nums))
//
// 	fmt.Println(caidai(8, 3, []int{1, 2, 3, 2, 1, 4, 5, 1}))
// }
