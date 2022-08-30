package golang

func isAnagram(s string, t string) bool {
	sList := stringList(s)
	tList := stringList(t)
	for k, v := range sList {
		if value, ok := tList[k]; !ok || v != value {
			return false
		}
	}
	return len(sList) == len(tList)
}

func stringList(s string) map[rune]uint {
	list := make(map[rune]uint)
	for _, v := range s {
		list[v]++
	}
	return list
}
