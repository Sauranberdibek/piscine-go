package piscine

func IsLower(str string) bool {
	r := []rune(str)
	ln := 0
	for i := range r {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if r[i] < 'a' || r[i] > 'z' {
			return false
		}
	}
	return true
}
