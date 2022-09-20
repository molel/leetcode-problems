package golang

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var combinations [][]int
	combinations = findCombinations(target, 0, 0, candidates, []int{}, combinations)
	return combinations
}

func findCombinations(target, currentSum, previousCandidate int, candidates, combination []int, combinations [][]int) [][]int {
	if previousCandidate <= len(candidates)-1 {
		for i := previousCandidate; i < len(candidates) && candidates[i]+currentSum <= target; i++ {
			if candidates[i]+currentSum == target {
				newCombination := append([]int{}, combination...)
				newCombination = append(newCombination, candidates[i])
				combinations = append(combinations, newCombination)
			} else {
				combinations = findCombinations(target, currentSum+candidates[i], i, candidates, append(combination, candidates[i]), combinations)
			}
		}
	}
	return combinations
}
