class Solution:
    def maxArea(self, heights: list[int]) -> int:
        left = 0
        right = len(heights) - 1
        maxVolume = (right - left) * min(heights[right], heights[left])
        while left < right:
            if heights[left] < heights[right]:
                left += 1
            else:
                right -= 1
            maxVolume = max(maxVolume, (right - left) * min(heights[right], heights[left]))
        return maxVolume