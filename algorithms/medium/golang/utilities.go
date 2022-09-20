package golang

func max(nums ...int) int {
	maxNum := -(1 << 31)
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func min(nums ...int) int {
	minNum := 1<<31 - 1
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}
