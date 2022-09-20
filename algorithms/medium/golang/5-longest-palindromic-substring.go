package golang

func longestPalindrome(s string) string {
	longestPalindrome := string(s[0])
	for i := 1; i < len(s); i++ {
		odd := maxPalindrome(s, i, i)
		even := maxPalindrome(s, i-1, i)
		if len(longestPalindrome) < len(odd) {
			longestPalindrome = odd
		}
		if len(longestPalindrome) < len(even) {
			longestPalindrome = even
		}
	}
	return longestPalindrome
}

func maxPalindrome(s string, left, right int) string {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
