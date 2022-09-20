package golang

func insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := make([][]int, 0)

	i := 0
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		newIntervals = append(newIntervals, intervals[i])
	}

	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}

	newIntervals = append(newIntervals, newInterval)

	for ; i < len(intervals); i++ {
		newIntervals = append(newIntervals, intervals[i])
	}

	return newIntervals
}
