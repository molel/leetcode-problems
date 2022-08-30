package golang

import "sort"

func change(amount int, coins []int) int {
	sort.Ints(coins)
	dp := make([][]int, 2)
	dp[1] = make([]int, amount+1)
	for i := len(coins) - 1; i >= 0; i-- {
		dp[0] = make([]int, amount+1)
		dp[0][0] = 1
		for amt := 1; amt <= amount; amt++ {
			if amt >= coins[i] {
				dp[0][amt] += dp[0][amt-coins[i]]
			}
			dp[0][amt] += dp[1][amt]
		}
		dp[1] = dp[0]
	}
	return dp[0][amount]
}
