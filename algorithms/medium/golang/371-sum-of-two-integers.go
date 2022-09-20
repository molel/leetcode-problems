package golang

func getSum(a int, b int) int {
	var res, carry int
	res = a ^ b
	carry = (a & b) << 1
	for carry != 0 {
		res, carry = res^carry, (res&carry)<<1
	}
	return res
}
