package golang

func majorityElement(nums []int) int {
	majorNum := nums[0]
	majorCount := 1
	for i := 1; i < len(nums); i++ {
		if majorCount == 0 {
			majorNum = nums[i]
			majorCount = 1
		} else if nums[i] == majorNum {
			majorCount += 1
		} else {
			majorCount -= 1
		}
	}
	return majorNum
}
