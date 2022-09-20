package golang

func numDecodings(s string) int {
	dp1 := 1
	dp2 := 0
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		dp := dp1
		if s[i] == '0' {
			dp = 0
		} else if i < n-1 && (s[i]-'0')*10+s[i+1]-'0' <= 26 {
			dp += dp2
		}
		dp2 = dp1
		dp1 = dp
	}
	return dp1
}
