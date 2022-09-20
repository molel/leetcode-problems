package golang

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	left := 0
	right := 1
	length := 1
	charIndexes := make(map[uint8]int)
	charIndexes[s[left]] = left
	for right < n {
		if index, ok := charIndexes[s[right]]; ok && index >= left {
			left = index + 1
		}
		charIndexes[s[right]] = right
		length = max(length, right-left+1)
		right++
	}
	return length
}
