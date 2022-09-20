package golang

func canJump(nums []int) bool {
	if nums[0] == 0 {
		return len(nums) == 1
	}
	maxIndex := 0
	lastIndex := len(nums) - 1
	for currentIndex, jumpLength := range nums[:lastIndex] {
		if currentIndex <= maxIndex {
			maxIndex = max(maxIndex, currentIndex+jumpLength)
		}
	}
	return maxIndex >= lastIndex
}
