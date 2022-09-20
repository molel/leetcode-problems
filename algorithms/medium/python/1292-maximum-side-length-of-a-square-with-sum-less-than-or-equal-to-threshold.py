class Solution:
    def maxSideLength(self, mat: List[List[int]], threshold: int) -> int:
        n, m = len(mat), len(mat[0])
        prefix_sum = [[0 for _ in range(m + 1)] for _ in range(n + 1)]
        side = 0
        for i in range(n):
            for j in range(m):
                prefix_sum[i + 1][j + 1] = mat[i][j] + prefix_sum[i][j + 1] + prefix_sum[i + 1][j] - prefix_sum[i][j]

                if side <= min(i, j) and prefix_sum[i + 1][j + 1] - prefix_sum[i - side][j + 1] - prefix_sum[i + 1][j - side] + prefix_sum[i - side][j - side] <= threshold:
                    side += 1
        return side

        