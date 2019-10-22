package piscine

func LastRune(s string) rune {
	c := 0
	for index := range s {
		c = index
	}
	rr := []rune(s)
	return rr[c]
}
