class Solution:
    def minFlips(self, a: int, b: int, c: int) -> int:
        flips = 0
        i = 0
        while (1 << i) <= max(a, b, c):
            if c & (1 << i):
                flips += (((a | b) & (1 << i)) >> i) ^ 1
            else:
                flips += ((a & (1 << i)) >> i) + ((b & (1 << i)) >> i)
            i += 1
        return flips
