package golang

func singleNumber(nums []int) int {
	var singleElement int
	for _, num := range nums {
		singleElement ^= num
	}
	return singleElement
}
