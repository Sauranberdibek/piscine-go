package piscine

func AlphaCount(str string) int {
	count := 0
	for _, let := range str {
		if let >= 'A' && let <= 'Z' || let >= 'a' && let <= 'z' {
			count++
		}
	}
	return count

}
