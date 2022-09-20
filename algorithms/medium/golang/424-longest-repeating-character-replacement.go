package golang

func characterReplacement(s string, k int) int {
	n := len(s)
	if n == k || n == 1 {
		return n
	}
	left := 0
	right := 0
	characterFrequency := make(map[byte]int)
	mostCommonCharacter := byte(0)
	length := 0
	for right < n {
		characterFrequency[s[right]]++
		if mostCommonCharacter == 0 || characterFrequency[mostCommonCharacter] < characterFrequency[s[right]] {
			mostCommonCharacter = s[right]
		}
		if right-left+1-characterFrequency[mostCommonCharacter] > k {
			characterFrequency[s[left]]--
			left++
		}
		length = max(length, right-left+1)
		right++
	}
	return length
}
