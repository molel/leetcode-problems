class Solution:
    def reverse(self, x: int) -> int:
        if not x:
            return x
        x = str(x)
        if x[0] == '-':
            x = int(x[0] + x[-1:0:-1])
        else:
            x = int(x[::-1])
        if x < -2 ** 31 or x > 2 ** 31 - 1:
            return 0
        return x