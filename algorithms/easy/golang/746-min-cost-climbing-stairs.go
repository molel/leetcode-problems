package golang

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, 2)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		newCost := cost[i] + min(dp[0], dp[1])
		dp[0], dp[1] = dp[1], newCost
	}
	return min(dp[0], dp[1])
}
