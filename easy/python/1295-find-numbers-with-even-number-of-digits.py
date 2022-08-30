class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        return len([0 for num in nums if int(log10(num)) % 2])