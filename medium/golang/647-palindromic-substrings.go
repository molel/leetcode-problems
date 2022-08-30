package golang

func countSubstrings(s string) int {
	numberOfPalindromes := 1
	for i := 1; i < len(s); i++ {
		countPalindromes(s, i, i, &numberOfPalindromes)
		countPalindromes(s, i-1, i, &numberOfPalindromes)
	}
	return numberOfPalindromes
}

func countPalindromes(s string, left, right int, numberOfPalindromes *int) {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		*numberOfPalindromes += 1
		left--
		right++
	}
}
