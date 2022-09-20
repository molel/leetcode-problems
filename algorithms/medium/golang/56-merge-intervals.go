package golang

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][]int, 1)
	mergedIntervals[0] = intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= mergedIntervals[len(mergedIntervals)-1][1] {
			mergedIntervals[len(mergedIntervals)-1][1] = max(intervals[i][1], mergedIntervals[len(mergedIntervals)-1][1])
		} else {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}

	return mergedIntervals
}
