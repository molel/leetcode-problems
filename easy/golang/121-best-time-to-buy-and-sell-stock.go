package golang

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	left := 0
	right := 1
	profit := prices[right] - prices[left]
	for right < n {
		if prices[right] < prices[left] {
			left = right
			right = left + 1
		} else {
			profit = max(profit, prices[right]-prices[left])
			right++
		}
	}
	return max(profit, 0)
}
