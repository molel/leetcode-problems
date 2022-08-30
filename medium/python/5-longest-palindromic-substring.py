class Solution:
    def longestPalindrome(self, s: str) -> str:
        longestPalindrome = s[0]
        for i in range(1, len(s)):
            oddPalindrome = self.getLongestPalindrome(s, i, i)
            evenPalindrome = self.getLongestPalindrome(s, i - 1, i)
            longestPalindrome = max(longestPalindrome, oddPalindrome, evenPalindrome, key=len)
        return longestPalindrome

    def getLongestPalindrome(self, s: str, left: int, right: int) -> str:
        while 0 <= left and right <= len(s) - 1:
            if s[left] != s[right]:
                break
            left -= 1
            right += 1
        return s[left + 1:right]