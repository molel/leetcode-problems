package golang

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	m := map[int]int{}
	for _, h := range hand {
		m[h]++
	}
	count := 0
	sort.Ints(hand)
	for _, h := range hand {
		if m[h] > 0 {
			count++
			for i := 0; i < groupSize; i++ {
				if m[h] <= 0 {
					return false
				} else {
					m[h]--
				}
				h++
			}
		}
	}
	if count == len(hand)/groupSize {
		return true
	}
	return false
}
