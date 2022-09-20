class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        LIS = [1 for _ in range(len(nums))]
        for i in range(len(nums) - 2, -1, -1):
            LIS[i] += max(max([LIS[j] for j in range(i + 1, len(nums)) if nums[i] < nums[j]], [0]))
        return max(LIS)
        