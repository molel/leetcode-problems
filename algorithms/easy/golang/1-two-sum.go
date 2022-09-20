package golang

func twoSum(nums []int, target int) []int {
	discoveredNums := make(map[int]int)
	for i, num := range nums {
		if j, ok := discoveredNums[target-num]; ok {
			return []int{j, i}
		}
		discoveredNums[num] = i
	}
	return []int{-1, -1}
}
