package golang

func titleToNumber(columnTitle string) int {
	var columnNumber int
	for _, char := range []byte(columnTitle) {
		columnNumber = columnNumber*26 + int(char-'A') + 1
	}
	return columnNumber
}
