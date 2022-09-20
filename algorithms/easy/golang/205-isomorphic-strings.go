package golang

func isIsomorphic(s string, t string) bool {
	n := len(s)
	sTable := map[byte]int{}
	sList := make([][]int, 0)
	tTable := map[byte]int{}
	tList := make([][]int, 0)
	for i := 0; i < n; i++ {
		if index, ok := sTable[s[i]]; !ok {
			sTable[s[i]] = len(sList)
			sList = append(sList, []int{i})
		} else {
			sList[index] = append(sList[index], i)
		}
		if index, ok := tTable[t[i]]; !ok {
			tTable[t[i]] = len(tList)
			tList = append(tList, []int{i})
		} else {
			tList[index] = append(tList[index], i)
		}
	}
	return isEqual(sList, tList)
}

func isEqual(first, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if len(first[i]) != len(second[i]) {
			return false
		}
		for j := 0; j < len(first[i]); j++ {
			if first[i][j] != second[i][j] {
				return false
			}
		}
	}
	return true
}
