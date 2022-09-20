package golang

func jump(nums []int) int {
	numberOfJumps := 0
	maxIndex := 0
	nextMaxIndex := 0
	lastIndex := len(nums) - 1
	for i := 0; maxIndex < lastIndex; i++ {
		nextMaxIndex = max(nextMaxIndex, i+nums[i])
		if i == maxIndex {
			maxIndex = nextMaxIndex
			numberOfJumps++
		}
	}
	return numberOfJumps
}
