package main

func maxProfit(prices []int, fee int) int {
	// buy in the first day
	var minBuy = prices[0]
	var res int
	for i := 0; i < len(prices); i++ {
		// if current price less than minBuy then buy immediately
		if prices[i] < minBuy {
			minBuy = prices[i]
		}
		//  if profit <= 0 then don't sale, find the next time to sale
		if prices[i] >= minBuy && prices[i]-fee-minBuy <= 0 {
			continue
		}
		// profit > 0 can sale
		if prices[i] > minBuy+fee {
			// accumulate everyday profit
			res += prices[i] - minBuy - fee
			// update minBuy
			minBuy = prices[i] - fee
		}
	}
	return res
}
