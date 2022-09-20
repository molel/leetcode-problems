class Solution:
    def longestPalindromeSubseq(self, s: str) -> int:
        length = len(s)
        dp = [[i == j for j in range(length)] for i in range(length)]
        for j in range(1, len(s)):
            for i in range(j - 1, -1, -1):
                if s[i] == s[j]:
                    dp[i][j] = 2 + dp[i + 1][j - 1]
                else:
                    dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
        return int(dp[0][-1])