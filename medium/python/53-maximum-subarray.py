class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        result = nums[0]
        sum = 0
        for i in range(0, len(nums)):
            sum += nums[i]
            result = max(result, sum)
            sum = max(sum, 0)
        return result
        