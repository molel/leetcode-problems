package golang

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for _, num := range nums {
		curSum += num
		maxSum = max(maxSum, curSum)
		curSum = max(curSum, 0)
	}
	return maxSum
}
