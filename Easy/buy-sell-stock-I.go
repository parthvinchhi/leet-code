package easy

func MaxProfit(prices []int) int {
	var maxProfit int
	var buy = prices[0]
	for key, val := range prices {
		if buy < prices[key] {
			profit := val - buy
			if maxProfit < profit {
				maxProfit = profit
			}
		} else {
			buy = prices[key]
		}
	}
	return maxProfit
}
