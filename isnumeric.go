package piscine

func IsNumeric(str string) bool {
	r := []rune(str)
	ln := 0
	for i := range r {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if r[i] < '0' || r[i] > '9' {
			return false
		}
	}
	return true
}
