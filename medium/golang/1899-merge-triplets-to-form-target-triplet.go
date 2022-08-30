package golang

func mergeTriplets(triplets [][]int, target []int) bool {
	validTriplets := make([][]int, 0, len(triplets))
	for _, triplet := range triplets {
		if isValid(triplet, target) {
			validTriplets = append(validTriplets, triplet)
		}
	}
	n := len(validTriplets)
	if n == 0 {
		return false
	}
	triplet := validTriplets[0]
	for i := 1; i < n; i++ {
		triplet = mergeTwoTriplets(triplet, validTriplets[i])
	}
	return isEqual(triplet, target)
}

func isValid(triplet, target []int) bool {
	return triplet[0] <= target[0] && triplet[1] <= target[1] && triplet[2] <= target[2]
}

func isEqual(firstTriplet, secondTriplet []int) bool {
	return firstTriplet[0] == secondTriplet[0] && firstTriplet[1] == secondTriplet[1] && firstTriplet[2] == secondTriplet[2]
}

func mergeTwoTriplets(firstTriplet, secondTriplet []int) []int {
	return []int{max(firstTriplet[0], secondTriplet[0]), max(firstTriplet[1], secondTriplet[1]), max(firstTriplet[2], secondTriplet[2])}
}
