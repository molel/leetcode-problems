package golang

func hammingWeight(num uint32) int {
	var bitsNumber int
	for num > 0 {
		bitsNumber += int(num & 1)
		num /= 2
	}
	return bitsNumber
}
