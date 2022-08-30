package golang

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	differences := make([]int, n)
	totalConsumption := 0
	for i := 0; i < n; i++ {
		differences[i] = gas[i] - cost[i]
		totalConsumption += differences[i]
	}
	if totalConsumption < 0 {
		return -1
	}
	start := 0
	current := 0
	consumption := 0
	length := 0
	for length < n {
		for consumption < 0 {
			consumption -= differences[start]
			start = (start + 1) % n
			length--
		}
		consumption += differences[current]
		current = (current + 1) % n
		length++
	}
	return start
}
