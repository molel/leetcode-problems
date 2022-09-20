class Solution:
    def countRectangles(self, rectangles: List[List[int]], points: List[List[int]]) -> List[int]:
        temp = dict()
        rectangles.sort()
        for length, height in rectangles:
            if height in temp:
                temp[height].append(length)
            else:
                temp.update({height: [length]})
        rectangles = temp
        heights = sorted(list(rectangles.keys()))
        count = [0] * len(points)
        for i, (x, y) in enumerate(points):
            min_height_i = self.search(heights, y)
            if min_height_i != -1:
                for height in heights[min_height_i:]:
                    min_length_i = self.search(rectangles[height], x)
                    if min_length_i != -1:
                        count[i] += len(rectangles[height]) - min_length_i
        return count

    def search(self, array, target):
            if target > array[-1]:
                return -1
            left, right = 0, len(array) - 1
            while left < right:
                mid = (right + left) // 2
                if target > array[mid]:
                    left = mid + 1
                else:
                    right = mid
            return right
