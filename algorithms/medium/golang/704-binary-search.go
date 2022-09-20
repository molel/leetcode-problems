package golang

func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	n := len(matrix)
	m := len(matrix[0])
	right := n*m - 1
	var mid int
	for right-left > 1 {
		mid = (left + right) / 2
		if matrix[mid/m][mid%m] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return matrix[mid/m][mid%m] == target
}
