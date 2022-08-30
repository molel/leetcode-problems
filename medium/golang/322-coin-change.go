package golang

func coinChange(coins []int, amount int) int {
	const MAX = 1<<16 - 1
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = MAX
	}
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}
	if dp[amount] == MAX {
		return -1
	}
	return dp[amount]
}
