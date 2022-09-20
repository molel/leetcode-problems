class Solution:
    def romanToInt(self, s: str) -> int:
        convert = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
        "E": 0
        }
        s += "E"
        result = 0
        i = 0
        while i < len(s) - 1:
            if convert[s[i + 1]] > convert[s[i]]:
                result += convert[s[i + 1]] - convert[s[i]]
                i += 2
            else:
                result += convert[s[i]]
                i += 1
        return result
