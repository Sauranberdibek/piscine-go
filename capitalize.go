package piscine

func Capitalize(s string) string {
	r := []rune(s)
	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] = rune(r[0] - 32)
	}

	len_a := 0

	for index := range r {
		len_a = index
	}

	for i := 1; i <= len_a; i++ {
		if (r[i-1] < 'A' || r[i-1] > 'Z') &&
			(r[i] >= 'a' && r[i] <= 'z') &&
			(r[i-1] < 'a' || r[i-1] > 'z') &&
			(r[i-1] < '0' || r[i-1] > '9') {
			r[i] = rune(r[i] - 32)
		}
		if (r[i-1] >= 'A' && r[i-1] <= 'Z') &&
			(r[i] >= 'A' && r[i] <= 'Z') {
			r[i] = rune(r[i] + 32)
		}
		if (r[i-1] >= 'a' && r[i-1] <= 'z') &&
			(r[i] >= 'A' && r[i] <= 'Z') {
			r[i] = rune(r[i] + 32)
		}
		if (r[i-1] >= '0' && r[i-1] <= '9') &&
			(r[i] >= 'A' && r[i] <= 'Z') {
			r[i] = rune(r[i] + 32)
		}

	}

	return string(r)
}
