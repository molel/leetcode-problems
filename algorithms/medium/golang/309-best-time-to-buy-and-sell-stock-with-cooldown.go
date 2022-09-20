package golang

import "math"

func maxProfit(prices []int) int {
	cooldown, sell, hold := 0, 0, math.MinInt64

	for _, price := range prices {

		prevCooldown, prevSell, prevHold := cooldown, sell, hold

		cooldown = max(prevCooldown, prevSell)

		sell = prevHold + price

		hold = max(prevHold, prevCooldown-price)

	}

	return max(sell, cooldown)
}
