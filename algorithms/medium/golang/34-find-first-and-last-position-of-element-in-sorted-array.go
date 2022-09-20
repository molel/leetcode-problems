package golang

import "sort"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || target < nums[0] || target > nums[n-1] {
		return []int{-1, -1}
	}
	indexes := []int{binarySearchLeft(nums, target), binarySearchRight(nums, target)}
	sort.Ints(indexes)
	return indexes
}

func binarySearchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for right-left > 1 {
		mid := (right + left) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	if nums[right] != target {
		return -1
	}
	return right
}

func binarySearchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for right-left > 1 {
		mid := (right + left) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
