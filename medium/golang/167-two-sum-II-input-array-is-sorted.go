package golang

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for currentSum := numbers[left] + numbers[right]; currentSum != target; currentSum = numbers[left] + numbers[right] {
		if currentSum > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}
