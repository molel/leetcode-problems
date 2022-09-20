class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        longest = 0
        long = 0
        queue = []
        for char in s:
            if char in queue:
                index = queue.index(char)
                queue[:index + 1] = []
                longest = max(long, longest)
                long -= index
                queue.append(char)
            else:
                queue.append(char)
                long += 1
        return max(long, longest)