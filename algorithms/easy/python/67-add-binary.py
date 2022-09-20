class Solution:
    def addBinary(self, a: str, b: str) -> str:
        remainder = 0
        result = []
        i = 1
        while i <= max(len(a), len(b)):
            sum_ = 0
            if i <= len(a):
                sum_ += int(a[-i])
            if i <= len(b):
                sum_ += int(b[-i])
            sum_ += remainder
            remainder = sum_ // 2
            sum_ %= 2
            result.append(str(sum_))
            i += 1
        if remainder:
            result.append(str(remainder))
        result.reverse()
        return "".join(result)