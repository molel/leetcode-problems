package golang

func reverseBits(num uint32) uint32 {
	var reverseNum uint32
	var bitNum = 31
	for bitNum >= 0 {
		reverseNum |= (num & 1) << bitNum
		num >>= 1
		bitNum--
	}
	return reverseNum
}
