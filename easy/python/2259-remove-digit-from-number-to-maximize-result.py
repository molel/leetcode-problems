class Solution:
    def removeDigit(self, number: str, digit: str) -> str:
        max_number = -1
        for i, el in enumerate(number):
            if el == digit:
                max_number = max(max_number, int(number[:i] + number[i + 1:]))
        return str(max_number)