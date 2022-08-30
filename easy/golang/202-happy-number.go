package golang

func isHappy(n int) bool {
	discoveredNumbers := make(map[int]bool)
	for sum := digitsSquaredSum(n); !discoveredNumbers[sum]; sum = digitsSquaredSum(sum) {
		discoveredNumbers[sum] = true
		if sum == 1 {
			return true
		}
	}
	return false
}

func digitsSquaredSum(n int) (sum int) {
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
