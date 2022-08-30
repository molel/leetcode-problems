class Solution:
    def mySqrt(self, x: int) -> int:
        if not x:
            return 0
        left = 1
        right = 2147483647
        while True:
            mid = left + (right - left) // 2
            if mid > x // mid:
                right = mid - 1
            elif mid + 1 > x // (mid + 1):
                return mid
            else:
                left = mid + 1