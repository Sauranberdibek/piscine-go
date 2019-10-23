package piscine

func IsUpper(str string) bool {
	r := []rune(str)
	ln := 0
	for i := range r {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if r[i] < 'A' || r[i] > 'Z' {
			return false
		}
	}
	return true
}
