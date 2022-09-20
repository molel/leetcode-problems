class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        nums.sort()
        triplets = set()
        for k in range(len(nums)):
            i = k + 1
            j = len(nums) - 1
            while i < j:
                if nums[k] + nums[i] + nums[j] == 0:
                    triplets.add((nums[k], nums[i], nums[j]))
                    i+=1
                elif nums[k] + nums[i] + nums[j] > 0:
                    j -= 1
                else:
                    i += 1
        return triplets
