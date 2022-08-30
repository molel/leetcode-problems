class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        result = []
        while nums1 or nums2:
            if not (nums1 and nums2):
                if nums1:
                    result.append(nums1.pop())
                elif nums2:
                    result.append(nums2.pop())
            else:
                if nums1[-1] >= nums2[-1]:
                    result.append(nums1.pop())
                else:
                    result.append(nums2.pop())
        result.reverse()
        if len(result) % 2:
            return result[len(result) // 2]
        else:
            return (result[len(result) // 2] + result[len(result) // 2 - 1]) / 2