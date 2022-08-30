package golang

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	substringFrequency := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		substringFrequency[s1[i]]++
	}

	matches := 0
	n := len(s1)

	for i := 0; i < len(s1); i++ {
		if frequency, ok := substringFrequency[s2[i]]; ok {
			if frequency > 0 {
				matches++
			}
			substringFrequency[s2[i]]--
		}
	}

	for i := n; matches < n && i < len(s2); i++ {
		if frequency, ok := substringFrequency[s2[i-n]]; ok {
			if frequency >= 0 {
				matches--
			}
			substringFrequency[s2[i-n]]++
		}
		if frequency, ok := substringFrequency[s2[i]]; ok {
			if frequency > 0 {
				matches++
			}
			substringFrequency[s2[i]]--
		}
	}

	return matches == n
}
