package golang

import "math"

func robII(nums []int) int {
	if length := len(nums); length == 1 {
		return nums[0]
	}
	n := len(nums)
	return int(math.Max(float64(robHouses(nums[0:n-1])), float64(robHouses(nums[1:n]))))
}

func robHouses(nums []int) int {
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
