class Solution:
    def fullBloomFlowers(self, flowers: List[List[int]], persons: List[int]) -> List[int]:
        starts, ends = zip(*flowers)
        starts = sorted(list(starts))
        ends = sorted(list(ends))
        return [self.binary_search_start(starts, i) - self.binary_search_end(ends, i) for i in persons]



    def binary_search_start(self, array, target):
        left, right = 0, len(array) - 1
        if target < array[0]:
            i = 0
        elif target >= array[-1]:
            i = len(array)
        else:
            while right - left > 1:
                mid = (right + left) // 2
                if target >= array[mid]:
                    left = mid
                else:
                    right = mid
            i = left + 1
        return i


    def binary_search_end(self, array, target):
        left, right = 0, len(array) - 1
        if target <= array[0]:
            i = 0
        elif target > array[-1]:
            i = len(array)
        else:
            while right - left > 1:
                mid = (right + left) // 2
                if target > array[mid]:
                    left = mid
                else:
                    right = mid
            i = left + 1
        return i