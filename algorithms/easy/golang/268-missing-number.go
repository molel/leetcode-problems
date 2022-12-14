package golang

func missingNumber(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		n = n ^ i ^ num
	}
	return n
}
