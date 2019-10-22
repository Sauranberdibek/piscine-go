package piscine

func NRune(s string, n int) rune {
	c := 0
	for index := range s {
		c = index
	}
	if (c+1) >= n && n > 0 {
		rr := []rune(s)
		return (rr[n-1])
	} else {
		return 0
	}
}
