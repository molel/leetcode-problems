package golang

func partitionLabels(s string) []int {
	partitions := make([]int, 0)
	lastIndexes := map[byte]int{}
	for i := 0; i < len(s); i++ {
		lastIndexes[s[i]] = i
	}
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		right = max(right, lastIndexes[s[i]])
		if i == right {
			partitions = append(partitions, right-left+1)
			left = right + 1
		}
	}
	return partitions
}
