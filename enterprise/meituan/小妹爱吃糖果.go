package main

/*描述：吃第i个糖果就不能吃第i-1, i-2 , i+1 ,i+2个，每个糖果有一个美味值，求最大美味值。*/

/*
dp[i]表示前i个糖果的最大美味值，
转移方程dp[i] = max(dp[i - 3] + a[i], dp[i - 1]) (i > 2)，
此外注意a.length分别等于0，1，2时的特殊情况处理
*/
func main() {

}

func candy(candys []int) int {
	dp := make([]int, len(candys)+1)
	dp[0] = candys[0]
	dp[1] = candyFindMax(candys[0], candys[1])
	dp[2] = candyFindMax(dp[1], candys[2])
	for i := 3; i <= len(candys); i++ {
		dp[i] = candyFindMax(dp[i-1], dp[i-3]+candys[i-1])
	}
	return dp[len(candys)]
}

func candyFindMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
