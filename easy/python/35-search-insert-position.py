class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                if nums[mid] == target and nums[mid - 1] != target:
                    return mid
                else:
                    right = mid - 1
        return left