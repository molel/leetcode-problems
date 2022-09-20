class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        triangle = [[1]]
        for i in range(2, numRows + 1):
            row = [1] + [0] * max(i - 2, 0) + [1]
            for j in range(1, i - 1):
                row[j] = triangle[-1][j - 1] + triangle[-1][j]
            triangle.append(row)
        return triangle