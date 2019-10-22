package piscine

func NRune(s string, n int) rune {
	ln := 0
	for i := range s {
		ln = i
	}

	tr := []rune(s)
	if ln >= n {
		return (tr[n-1])
	} else {
		return 0
	}
}
