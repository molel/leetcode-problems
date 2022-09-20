class Solution:
    def summaryRanges(self, nums: list[int]) -> list[str]:
        if not nums:
            return []
        ranges = []
        start = end = nums[0]
        for i in range(1, len(nums)):
            if nums[i] - nums[i - 1] == 1:
                end = nums[i]
            else:
                ranges.append(self.makeRange(start, end))
                start = end = nums[i]
        ranges.append(self.makeRange(start, end))
        return ranges

    def makeRange(self, start: int, end: int) -> str:
        if start == end:
            return str(start)
        else:
            return f"{start}->{end}"