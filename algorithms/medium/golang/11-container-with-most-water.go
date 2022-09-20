package golang

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := (right - left) * min(height[right], height[left])
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		area = max(area, (right-left)*min(height[right], height[left]))
	}
	return area
}
