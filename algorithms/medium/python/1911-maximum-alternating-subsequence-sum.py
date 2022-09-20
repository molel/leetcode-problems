class Solution:
    def maxAlternatingSum(self, nums: List[int]) -> int:
        even = odd = 0
        for element in nums:
            even, odd = max(even, odd + element), max(odd, even - element)
        return even