class Solution:
    def countLatticePoints(self, circles: List[List[int]]) -> int:
        points = set()
        for circle in circles:
            for x in range(circle[0] - circle[2], circle[0] + circle[2] + 1):
                for y in range(circle[1] - circle[2], circle[1] + circle[2] + 1):
                    if (x - circle[0]) * (x - circle[0]) + (y - circle[1]) * (y - circle[1]) <= circle[2] ** 2:
                        points.add((x, y))
        return len(points)