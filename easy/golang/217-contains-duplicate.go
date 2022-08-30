package golang

func containsDuplicate(nums []int) bool {
	numberAppearance := make(map[int]bool)
	for _, num := range nums {
		if numberAppearance[num] {
			return true
		}
		numberAppearance[num] = true
	}
	return false
}
