package piscine

func IsAlpha(str string) bool {
	r := []rune(str)
	k := true
	len := 0

	for i := range r {
		len = i
	}

	for i := 0; i <= len; i++ {
		if (r[i] < 'A' || r[i] > 'Z') &&
			(r[i] < 'a' || r[i] > 'z') &&
			(r[i] < '0' || r[i] > '9') {
			k = false
			break
		}
	}
	return k
}
