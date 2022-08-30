package golang

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern, _ := regexp.Compile(`[a-z0-9]`)
	slice := pattern.FindAllString(strings.ToLower(s), -1)
	left := 0
	right := len(slice) - 1
	for left < right {
		if slice[left] != slice[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
