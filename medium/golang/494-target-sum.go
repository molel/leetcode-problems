package golang

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[[2]int]int)

	var subFunc func(int, int) int
	subFunc = func(index, total int) int {
		if index == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}
		if numOfCombinations, ok := dp[[2]int{index, total}]; ok {
			return numOfCombinations
		}
		dp[[2]int{index, total}] = subFunc(index+1, total+nums[index]) + subFunc(index+1, total-nums[index])
		return dp[[2]int{index, total}]
	}

	return subFunc(0, 0)
}
