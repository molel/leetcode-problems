package golang

func maxProduct(nums []int) int {
	maxProduct := 1
	minProduct := 1
	product := nums[0]
	for _, num := range nums {
		if num > product {
			product = num
		}
	}
	for _, num := range nums {
		if num != 0 {
			tempMaxProduct := maxProduct
			maxProduct = max(maxProduct*num, minProduct*num, num)
			minProduct = min(tempMaxProduct*num, minProduct*num, num)
			product = max(product, maxProduct)
		} else {
			maxProduct = 1
			minProduct = 1
		}
	}
	return product
}
