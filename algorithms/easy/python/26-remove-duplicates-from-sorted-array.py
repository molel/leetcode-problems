class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i = 1
        previous = 0
        while i < len(nums):
            if nums[i] == nums[previous]:
                nums[i] = "_"
            else:
                previous += 1
                nums[i], nums[previous] = nums[previous], nums[i]
            i += 1
        return previous + 1
