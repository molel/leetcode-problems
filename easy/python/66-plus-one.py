class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits[-1] += 1
        i = len(digits) - 1
        while i >= 0:
            if digits[i] > 9:
                if i > 0:
                    digits[i - 1] += 1
                    digits[i] %= 10
                    i -= 1
                else:
                    digits.insert(0, 1)
                    digits[i + 1] %= 10
            else:
                return digits
        return digits