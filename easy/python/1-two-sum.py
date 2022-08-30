class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = dict()
        for i in range(len(nums)):
            element_to_find = target - nums[i]
            if element_to_find in hashmap:
                return [i, hashmap[element_to_find]]
            hashmap[nums[i]] = i
