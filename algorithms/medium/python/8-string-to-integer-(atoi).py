class Solution:
    def myAtoi(self, s: str) -> int:
        s = s.lstrip()
        if not s:
            return 0
        if s[0] in '-+':
            sign = s[0]
            start = 1
        else:
            sign = ''
            start = 0
        result = 0
        for char in s[start:]:
            if char.isnumeric():
                result = result * 10 + int(char)
            else:
                break
        if sign == '-':
            if result > 1 << 31:
                result = (1 << 31)
            result *= -1
        elif result > (1 << 31) - 1:
            result = (1 << 31) - 1
        return result