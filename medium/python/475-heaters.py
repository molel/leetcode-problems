class Solution:
    def findRadius(self, houses: list[int], heaters: list[int]) -> int:
        heaters.sort()
        radius = -1
        for house in houses:
            i, j = self.binarySearch(heaters, house)
            radius = max(radius, min(abs(heaters[i] - house), abs(heaters[j] - house)))
        return radius

    def binarySearch(self, array: list[int], target: int) -> tuple[int, int]:
        left, right = 0, len(array) - 1
        while right - left > 1:
            mid = (left + right) // 2
            if target == array[mid]:
                return mid, mid
            elif target > array[mid]:
                left = mid
            else:
                right = mid
        return left, right