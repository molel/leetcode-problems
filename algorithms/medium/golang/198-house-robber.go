package golang

import "math"

func rob(nums []int) int {
	if length := len(nums); length == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = int(math.Max(float64(dp[i]), float64(nums[i]+dp[i-1])))
	}
	return dp[len(nums)]
}
