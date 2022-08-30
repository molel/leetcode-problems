package golang

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	triplets := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		left := i + 1
		right := n - 1
		for left < right {
			if sum := nums[left] + nums[right]; sum == target {
				triplets = append(triplets, []int{-target, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return triplets
}
