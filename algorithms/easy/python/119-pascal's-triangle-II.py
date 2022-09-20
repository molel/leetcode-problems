class Solution:
    from math import factorial
    def getRow(self, rowIndex: int) -> List[int]:
        row = []

        for i in range(rowIndex + 1):
            row.append(self.combinations(rowIndex, i))

        return row

    def combinations(self, n: int, k: int) -> int:
        return round(factorial(n) / factorial(k) / factorial(n - k))