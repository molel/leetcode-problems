class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        prefix = ""
        i = 0
        while i <= 200:
            if i >= len(strs[0]):
                return prefix
            letter = strs[0][i]
            for word in strs:
                if i >= len(word) or word[i] != letter:
                    return prefix
            prefix += letter
            i += 1